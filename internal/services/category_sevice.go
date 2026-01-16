package services

import (
	"api-perpus-go/internal/models"
	"api-perpus-go/internal/repositories"
	"errors"
)

func CreateCategory(name string) (*models.Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	c := &models.Category{Name: name}
	if err := repositories.CreateCategory(c); err != nil {
		return nil, err
	}
	return c, nil
}

func ListCategories() ([]models.Category, error) {
	return repositories.GetAllCategories()
}

func GetCategory(id int) (*models.Category, error) {
	return repositories.GetCategoryByID(id)
}

func UpdateCategory(id int, name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return repositories.UpdateCategory(id, name)
}

func DeleteCategory(id int) error {
	return repositories.DeleteCategory(id)
}
