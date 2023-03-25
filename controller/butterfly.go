package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/easilok/lymantria-api/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm/clause"
)

// GET /butterfly
// Get registered butterflies
func (h *BaseHandler) GetButterfly(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}
	var butterflies []models.Butterfly
	h.db.Model(&models.Butterfly{}).Where("butterflies.deleted_at IS NULL").Preload("Details").Find(&butterflies)

	c.JSON(http.StatusOK, gin.H{"records": butterflies})
}

// PUT /butterfly/:butterflyId
// Update a butterfly
func (h *BaseHandler) UpdateButterfly(c *gin.Context) {

	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}
	// Validate input
	var input models.Butterfly
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Failed to parse input: %s", err.Error())},
		)
		return
	}

	butterflyId := c.Param("butterflyId")
	var editingButterfly models.Butterfly
	var err error
	if err = h.db.Where("id = ?", butterflyId).First(&editingButterfly).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Updates current butterfly details
	if err = h.db.Model(&editingButterfly).Updates(input).Error; err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Failed to update butterfly: %s", err.Error())},
		)
	}

	var editedButterfly models.Butterfly
	if err = h.db.Where("id = ?", butterflyId).First(&editedButterfly).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"record": editedButterfly})
}

// DELETE /butterfly/:butterflyId
// Delete a butterfly
func (h *BaseHandler) DeleteButterfly(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}
	butterflyId := c.Param("butterflyId")

	// if filename exists on storage -> delete it -> remove from note information
	var deletingButterfly models.Butterfly
	if err := h.db.Where("id = ?", butterflyId).First(&deletingButterfly).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := h.db.Model(&deletingButterfly).Update("deleted_at", time.Now()).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete butterfly"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "message": "Butterfly deleted"})
}
