package dtos

import "time"

type ProductCreateRequest struct {
	Name       string  `json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required,gt=0"`
	CategoryID string  `json:"category_id" validate:"required,uuid"`
}

type ProductUpdateRequest struct {
	ID         string  `validate:"required"`
	Name       string  `json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required,gt=0"`
	CategoryID string  `json:"category_id" validate:"required,uuid"`
}

type ProductResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Price     float64           `json:"price"`
	Category  *CategoryResponse `json:"category"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
