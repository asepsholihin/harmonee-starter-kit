package seeders

import (
	"harmonee/app/database/fakers"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	seeders := []Seeder{}
	for range 30 {
		seeders = append(seeders, Seeder{Seeder: fakers.PostFaker(db)})
		seeders = append(seeders, Seeder{Seeder: fakers.CategoryFaker(db)})
	}
	return seeders
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
