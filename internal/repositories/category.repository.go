package repositories

import (
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/model/dtos"
	"go-fiber-project-template/internal/model/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(request *dtos.CategorySearchRequest) ([]entities.Category, int64, error)
	FindById(categoryId string) (*entities.Category, error)
	CreateData(category *entities.Category) (*entities.Category, error)
	UpdateData(category *entities.Category) (*entities.Category, error)
	DeleteData(category *entities.Category) error
	FilterData(request *dtos.CategorySearchRequest) func(tx *gorm.DB) *gorm.DB
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

func (r *CategoryRepositoryImpl) FindAll(request *dtos.CategorySearchRequest) ([]entities.Category, int64, error) {
	var categories []entities.Category
	if err := r.GlobalRepositoryImpl.DB.Scopes(r.FilterData(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Order("created_at DESC").Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := r.GlobalRepositoryImpl.DB.Model(&entities.Category{}).Scopes(r.FilterData(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *CategoryRepositoryImpl) FindById(categoryId string) (*entities.Category, error) {
	var category *entities.Category

	err := r.GlobalRepositoryImpl.DB.First(&category, "id = ?", categoryId).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return category, nil
}

func (r *CategoryRepositoryImpl) FilterData(request *dtos.CategorySearchRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}

		return tx
	}
}
