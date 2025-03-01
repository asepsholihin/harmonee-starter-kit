package fakers

import (
	"harmonee/app/models"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func PostFaker(db *gorm.DB) *models.Post {
	user := UserFaker(db)
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	title := faker.Sentence()
	return &models.Post{
		UserID:         user.ID,
		Title:          title,
		Slug:           slug.Make(title),
		SummaryContent: faker.Paragraph(),
		Content:        faker.Paragraph(),
		Viewers:        0,
		PublishStatus:  2,
		PublishedAt:    time.Now(),
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		DeletedAt:      gorm.DeletedAt{},
		CreatedBy:      user.ID,
		UpdatedBy:      user.ID,
	}
}
