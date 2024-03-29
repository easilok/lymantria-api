package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/easilok/lymantria-api/database"
	"github.com/easilok/lymantria-api/helpers"
	"github.com/easilok/lymantria-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserPassword struct {
	Password string `json:"password"`
}

//A sample use
// var user = User{
// 	ID:       1,
// 	Username: "username",
// 	Password: "password",
// }

// POST /login
func (h *BaseHandler) Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid user provided")
		return
	}
	var logginUser models.User
	if err := h.db.Where("email = ?", u.Username).First(&logginUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// fmt.Printf("Received user: %v \nFetched user: %v\n", u, logginUser)
	//compare the user from the request, with the one we defined:
	if passwordOk := models.CheckPasswordHash(u.Password, logginUser.Password); !passwordOk {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := helpers.CreateToken(uint64(logginUser.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if database.EnableRedis {
		saveErr := database.CreateAuth(uint64(logginUser.ID), ts)
		if saveErr != nil {
			c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		}
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

// GET /Logout
func (h *BaseHandler) Logout(c *gin.Context) {
	au, err := helpers.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	if database.EnableRedis {
		deleted, delErr := database.DeleteAuth(au.AccessUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

// POST /Refresh
func (h *BaseHandler) Refresh(c *gin.Context) {
	if !database.EnableRedis {
		c.JSON(http.StatusUnauthorized, nil)
	}
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := database.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := helpers.CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := database.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}

// POST /register
func (h *BaseHandler) Register(c *gin.Context) {
	var u RegisterInput
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid user provided")
		return
	}
	var existingUser models.User
	if err := h.db.Where("email = ?", u.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	var newUser models.User

	hashedPassword, err := models.HashPassword(u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}

	newUser.Email = u.Username
	newUser.Name = u.Name
	newUser.Password = hashedPassword
	h.db.Save(&newUser)

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

// PATCH /password
func (h *BaseHandler) Password(c *gin.Context) {
	var u UserPassword
	// Must be logged in
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}
	if err := c.ShouldBindJSON(&u); err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, "Invalid user provided")
		return
	}

	var existingUser models.User
	if err := h.db.Where("id = ?", userId).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	hashedPassword, err := models.HashPassword(u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}

	existingUser.Password = hashedPassword
	h.db.Save(&existingUser)

	c.JSON(http.StatusOK, gin.H{"message": "User password updated"})
}
