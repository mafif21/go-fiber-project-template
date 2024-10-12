package repositories

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/model/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(productId string) (*entities.Product, error)
	CreateData(product *entities.Product) (*entities.Product, error)
	UpdateData(product *entities.Product) (*entities.Product, error)
	DeleteData(product *entities.Product) error
}

type ProductRepositoryImpl struct {
	GlobalRepositoryImpl[entities.Product]
	Log *logrus.Logger
}

func NewProductRepositoryImpl(db *gorm.DB, log *logrus.Logger) ProductRepository {
	return &ProductRepositoryImpl{
		GlobalRepositoryImpl: GlobalRepositoryImpl[entities.Product]{
			DB: db,
		},
		Log: log,
	}
}

func (r *ProductRepositoryImpl) FindAll() ([]entities.Product, error) {
	var products []entities.Product

	err := r.GlobalRepositoryImpl.DB.Find(&products).Error
	if err != nil {
		return nil, errors.New("invalid database")
	}

	return products, nil
}

func (r *ProductRepositoryImpl) FindById(productId string) (*entities.Product, error) {
	var product *entities.Product

	err := r.GlobalRepositoryImpl.DB.First(&product, "id = ?", productId).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return product, nil
}
