package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"
)

func GetAllCategories() ([]dto.CategoryDTO, error) {
	var categories []models.Categories
	var categoryDTOs []dto.CategoryDTO

	resp := database.DB.Find(&categories)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, dto.CategoryDTO{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categoryDTOs, nil
}

func GetCategoryById(categoryID string) (*dto.CategoryDTO, error) {
	var category models.Categories

	resp := database.DB.First(&category, "id = ?", categoryID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &dto.CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func CreateCategory(categoryDTO dto.CategoryCreateDTO) error {
	category := models.Categories{
		Name: categoryDTO.Name,
	}

	return database.DB.Create(&category).Error
}

func UpdateCategory(categoryID string, categoryDTO dto.CategoryUpdateDTO) error {
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return err
	}

	editCategory := models.Categories{
		ID:   uint(id),
		Name: categoryDTO.Name,
	}

	return database.DB.Updates(&editCategory).Error
}

func DeleteCategory(categoryID string) error {
	return database.DB.Delete(&models.Categories{}, "id = ?", categoryID).Error
}
