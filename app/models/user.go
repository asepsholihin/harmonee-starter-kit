package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type NullTime struct {
	sql.NullTime
}

type User struct {
	ID              uint           `json:"id" gorm:"primary_key;autoIncrement"`
	Name            string         `json:"name" gorm:"size:255;not null"`
	Username        string         `json:"username" gorm:"size:255;not null"`
	Password        string         `json:"-" gorm:"size:255;not null"`
	RememberToken   string         `json:"-" gorm:"size:255;not null"`
	Email           string         `json:"email" gorm:"size:255;not null"`
	EmailVerifiedAt NullTime       `json:"email_verified_at"`
	Phone           string         `json:"phone" gorm:"size:255;not null"`
	Status          bool           `json:"status"`
	AvatarIcon      string         `json:"avatar_icon" gorm:"size:255;not null"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`
	DeletedBy       uint           `json:"-"`
}

type AttributeUser struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Status     bool   `json:"status"`
	AvatarIcon string `json:"avatar_icon"`
}

func (r NullTime) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.Time)
	} else {
		return json.Marshal(nil)
	}
}
