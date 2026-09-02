package dto

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryCreateDTO struct {
	Name string `json:"name"`
}

type CategoryUpdateDTO struct {
	Name string `json:"name"`
}
