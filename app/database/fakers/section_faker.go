package fakers

import (
	"harmonee/app/models"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func SectionFaker(db *gorm.DB) *models.Section {
	name := faker.Sentence()
	return &models.Section{
		Name:        name,
		Slug:        slug.Make(name),
		Description: faker.Paragraph(),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		CreatedBy:   1,
		UpdatedBy:   1,
	}
}
