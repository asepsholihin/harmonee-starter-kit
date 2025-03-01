package models

import (
	"time"

	"gorm.io/gorm"
)

/* PUBLISH STATUS VALUE
1 = Draft
2 = Publish
3 = Archived
*/

type Post struct {
	ID             uint `json:"id" gorm:"primary_key;autoIncrement"`
	UserID         uint `json:"user_id"`
	User           User
	Categories     []Category     `gorm:"many2many:post_categories;"`
	ImageLarge     string         `json:"image_large" gorm:"size:255;"`
	ImageMedium    string         `json:"image_medium" gorm:"size:255;"`
	ImageSmall     string         `json:"image_small" gorm:"size:255;"`
	Title          string         `json:"title" gorm:"size:255;not null"`
	Slug           string         `json:"slug" gorm:"size:255;not null"`
	SummaryContent string         `json:"summary_content" gorm:"type:text"`
	Content        string         `json:"content" gorm:"type:text"`
	Viewers        uint           `json:"viewers"`
	PublishStatus  int8           `json:"publish_status"`
	PublishedAt    time.Time      `json:"published_at"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-"`
	CreatedBy      uint           `json:"-"`
	UpdatedBy      uint           `json:"-"`
	DeletedBy      uint           `json:"-"`
}
