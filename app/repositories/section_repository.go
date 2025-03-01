package repositories

import (
	"harmonee/app/database"
	"harmonee/app/models"
)

type SectionRepository interface {
	FindAll(perPage int, page int) ([]models.Section, int64, error)
	FindById(id int) (models.Section, error)
	Store(section models.Section) error
	Update(attr models.AttrSectionPost) error
	Delete(attr models.AttrSectionPost) error
}

type SectionRepositoryImplement struct{}

// FindAll implements
func (*SectionRepositoryImplement) FindAll(perPage int, page int) ([]models.Section, int64, error) {
	var section models.Section
	var sections []models.Section
	var count int64
	offset := (page - 1) * perPage

	// Total rows
	err := database.Connect.Model(&section).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	// Get data
	err = database.Connect.Model(&section).Order("created_at desc").Limit(perPage).Offset(offset).Find(&sections).Error
	if err != nil {
		return nil, 0, err
	}
	return sections, count, nil
}

// FindById implements
func (*SectionRepositoryImplement) FindById(id int) (models.Section, error) {
	var section models.Section

	// Get data
	err := database.Connect.Model(&section).Find(&section, id).Error
	if err != nil {
		return section, err
	}
	return section, nil
}

// Store implements
func (*SectionRepositoryImplement) Store(section models.Section) error {
	// Update data
	err := database.Connect.Model(&section).Create(&section).Error
	if err != nil {
		return err
	}
	return nil
}

// Update implements
func (*SectionRepositoryImplement) Update(attr models.AttrSectionPost) error {
	var section models.Section
	// Update data
	err := database.Connect.Model(&section).Where("id = ?", attr.ID).Updates(&attr).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete implements
func (*SectionRepositoryImplement) Delete(attr models.AttrSectionPost) error {
	var section models.Section
	// Delete data
	err := database.Connect.Model(&section).Where("id = ?", attr.ID).Delete(&section).Error
	if err != nil {
		return err
	}
	return nil
}

func NewSectionRepository() SectionRepository {
	return &SectionRepositoryImplement{}
}
