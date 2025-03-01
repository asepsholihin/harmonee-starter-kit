package services

import (
	"harmonee/app/models"
	"harmonee/app/repositories"
	"log"
)

type SectionService interface {
	FindAll(perPage int, page int) ([]models.Section, int64, error)
	FindById(id int) (models.Section, error)
	Store(section models.Section) error
	Update(attr models.AttrSectionPost) error
	Delete(attr models.AttrSectionPost) error
}

type SectionServiceImplement struct {
	sectionRepo repositories.SectionRepository
}

func (ss *SectionServiceImplement) FindAll(perPage int, page int) ([]models.Section, int64, error) {

	sections, count, err := ss.sectionRepo.FindAll(perPage, page)
	if err != nil {
		log.Fatal(err.Error())
		return sections, count, err
	}
	return sections, count, nil
}

func (ss *SectionServiceImplement) FindById(id int) (models.Section, error) {

	section, err := ss.sectionRepo.FindById(id)
	if err != nil {
		log.Fatal(err.Error())
		return section, err
	}
	return section, nil
}

func (ss *SectionServiceImplement) Store(section models.Section) error {

	err := ss.sectionRepo.Store(section)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (ss *SectionServiceImplement) Update(attr models.AttrSectionPost) error {

	err := ss.sectionRepo.Update(attr)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (ss *SectionServiceImplement) Delete(attr models.AttrSectionPost) error {

	err := ss.sectionRepo.Delete(attr)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func NewSectionService(sectionRepo *repositories.SectionRepository) SectionService {
	return &SectionServiceImplement{
		sectionRepo: *sectionRepo,
	}
}
