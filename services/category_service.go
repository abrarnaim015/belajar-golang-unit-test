package services

import (
	"errors"

	"github.com/abrarnaim015/belajar-golang-unit-test/entity"
	"github.com/abrarnaim015/belajar-golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error)  {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}