package fakers

import (
	"harmonee/app/models"
	"time"

	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	var password = string(hashPassword)

	return &models.User{
		Name:      faker.Name(),
		Username:  faker.Username(),
		Password:  password,
		Email:     faker.Email(),
		Phone:     faker.Phonenumber(),
		Status:    true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
