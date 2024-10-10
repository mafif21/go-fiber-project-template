package converters

import (
	"go-fiber-project-template/internal/model/dtos"
	"go-fiber-project-template/internal/model/entities"
)

func CategoryToResponse(category *entities.Category) *dtos.CategoryResponse {
	return &dtos.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
