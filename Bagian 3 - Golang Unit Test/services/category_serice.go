package services

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CatgoryService struct {
	Repository repository.CategoryRepository
}

func (service CatgoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}
}
