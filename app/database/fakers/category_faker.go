package fakers

import (
	"harmonee/app/models"
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func CategoryFaker(db *gorm.DB) *models.Category {
	section := SectionFaker(db)
	err := db.Create(&section).Error
	if err != nil {
		log.Fatal(err)
	}
	title := faker.Sentence()
	return &models.Category{
		SectionID:   section.ID,
		Name:        title,
		Slug:        slug.Make(title),
		Description: faker.Paragraph(),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		CreatedBy:   1,
		UpdatedBy:   1,
	}
}
