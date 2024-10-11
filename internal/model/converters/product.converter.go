package converters

import (
	"go-fiber-project-template/internal/model/dtos"
	"go-fiber-project-template/internal/model/entities"
)

func ProductToResponse(product *entities.Product) *dtos.ProductResponse {
	return &dtos.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Category:  CategoryToResponse(product.Category),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
