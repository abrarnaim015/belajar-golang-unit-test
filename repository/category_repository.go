package repository

import "github.com/abrarnaim015/belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}