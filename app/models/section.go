package models

import (
	"time"

	"gorm.io/gorm"
)

type Section struct {
	ID          uint           `json:"id" gorm:"primary_key;autoIncrement"`
	Name        string         `json:"name" gorm:"size:255" validate:"required"`
	Slug        string         `json:"slug" gorm:"size:255"`
	Description string         `json:"description" gorm:"type:text" validate:"required"`
	Icon        string         `json:"icon" gorm:"type:text"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	CreatedBy   uint           `json:"-"`
	UpdatedBy   uint           `json:"-"`
	DeletedBy   uint           `json:"-"`
	Categories  []Category
}

type AttrSectionPost struct {
	ID          uint
	Name        string
	Description string
}

type AttrCreateSectionPost struct {
	Name        string
	Description string
}
