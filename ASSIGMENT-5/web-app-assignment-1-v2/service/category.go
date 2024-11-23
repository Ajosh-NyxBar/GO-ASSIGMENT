package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryService struct {
	categoryRepo repo.CategoryRepository
}

func NewCategoryService(categoryRepo repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) Store(category *model.Category) error {
	err := s.categoryRepo.Store(category)
	if err != nil {
		return err
	}

	return nil
}

func (s *categoryService) Update(id int, category model.Category) error {
	return s.categoryRepo.Update(id, category)
}

func (s *categoryService) Delete(id int) error {
	return s.categoryRepo.Delete(id)
}

func (s *categoryService) GetByID(id int) (*model.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetList() ([]model.Category, error) {
	categories, err := s.categoryRepo.GetList()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
