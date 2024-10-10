package repositories

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/model/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]entities.Category, error)
	FindById(categoryId string) (*entities.Category, error)
	CreateData(category *entities.Category) (*entities.Category, error)
	UpdateData(category *entities.Category) (*entities.Category, error)
	DeleteData(category *entities.Category) error
}

type CategoryRepositoryImpl struct {
	GlobalRepositoryImpl[entities.Category]
	Log *logrus.Logger
}

func NewCategoryRepositoryImpl(db *gorm.DB, log *logrus.Logger) CategoryRepository {
	return &CategoryRepositoryImpl{
		GlobalRepositoryImpl: GlobalRepositoryImpl[entities.Category]{
			DB: db,
		},
		Log: log,
	}
}

func (r *CategoryRepositoryImpl) FindAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.GlobalRepositoryImpl.DB.Find(&categories).Error

	if err != nil {
		return nil, errors.New("invalid database")
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) FindById(categoryId string) (*entities.Category, error) {
	var category *entities.Category

	err := r.GlobalRepositoryImpl.DB.First(&category, "id = ?", categoryId).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return category, nil
}
