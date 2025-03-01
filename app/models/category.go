package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint `json:"id" gorm:"primary_key;autoIncrement"`
	ParentID    uint `json:"parent_id"`
	Section     Section
	SectionID   uint           `json:"section_id"`
	Posts       []Post         `gorm:"many2many:post_categories;"`
	Name        string         `json:"name" gorm:"size:255"`
	Slug        string         `json:"slug" gorm:"size:255"`
	Description string         `json:"description" gorm:"type:text"`
	Icon        string         `json:"icon" gorm:"type:text"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	CreatedBy   uint           `json:"-"`
	UpdatedBy   uint           `json:"-"`
	DeletedBy   uint           `json:"-"`
}
