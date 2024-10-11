package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID         string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string    `gorm:"unique;column:name;not null"`
	Price      float64   `gorm:"column:price;not null;check:price >= 0"`
	CategoryID string    `gorm:"type:uuid;not null"`
	Category   *Category `gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Product) TableName() string {
	return "products"
}

func (c *Product) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.NewString()
	return nil
}
