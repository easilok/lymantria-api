package models

import (
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	Email       string         `json:"email" gorm:"uniqueIndex;not null"`
	Password    string         `json:"password" gorm:"not null"`
	Name        string         `json:"name" gorm:"not null"`
	Permissions pq.StringArray `json:"permissions" gorm:"not null;type:text[];default:ARRAY[]::text[]" pg:",array"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
