package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type GlobalRepositoryImpl[T any] struct {
	DB *gorm.DB
}

func (r GlobalRepositoryImpl[T]) CreateData(entity *T) (*T, error) {
	if err := r.DB.Create(&entity).Error; err != nil {
		return nil, errors.New("failed to save data")
	}

	return entity, nil
}

func (r GlobalRepositoryImpl[T]) UpdateData(entity *T) (*T, error) {
	if err := r.DB.Updates(&entity).Error; err != nil {
		return nil, errors.New("failed to update data")
	}
	return entity, nil
}

func (r GlobalRepositoryImpl[T]) DeleteData(entity *T) error {
	if err := r.DB.Delete(&entity).Error; err != nil {
		return errors.New("failed to delete data")
	}

	return nil
}
