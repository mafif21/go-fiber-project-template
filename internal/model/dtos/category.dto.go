package dtos

import "time"

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategorySearchRequest struct {
	Name string `json:"name" validate:"max=100"`
	*SearchConf
}
