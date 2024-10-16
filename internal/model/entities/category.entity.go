package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `gorm:"unique;column:name;not null"`
	Products  []Product `gorm:"foreignKey:CategoryID;references:ID"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.NewString()
	return nil
}
