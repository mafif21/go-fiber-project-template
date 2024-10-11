package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/model/converters"
	"go-fiber-project-template/internal/model/dtos"
	"go-fiber-project-template/internal/model/entities"
	"go-fiber-project-template/internal/repositories"
)

type CategoryService interface {
	Get() ([]dtos.CategoryResponse, error)
	GetById(categoryId string) (*dtos.CategoryResponse, error)
	Create(request *dtos.CategoryCreateRequest) (*dtos.CategoryResponse, error)
	Update(request *dtos.CategoryUpdateRequest) (*dtos.CategoryResponse, error)
	Delete(categoryId string) error
}

type CategoryServiceImpl struct {
	Validate           *validator.Validate
	Log                *logrus.Logger
	CategoryRepository repositories.CategoryRepository
}

func NewCategoryServiceImpl(validate *validator.Validate, categoryRepository repositories.CategoryRepository, log *logrus.Logger) CategoryService {
	return &CategoryServiceImpl{
		Validate:           validate,
		CategoryRepository: categoryRepository,
		Log:                log,
	}
}

func (s CategoryServiceImpl) Get() ([]dtos.CategoryResponse, error) {
	datas, err := s.CategoryRepository.FindAll()
	if err != nil {
		s.Log.Warnf("failed get all data categories : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := make([]dtos.CategoryResponse, len(datas))
	for i, data := range datas {
		response[i] = *converters.CategoryToResponse(&data)
	}

	return response, nil
}

func (s CategoryServiceImpl) GetById(categoryId string) (*dtos.CategoryResponse, error) {
	category, err := s.CategoryRepository.FindById(categoryId)
	if err != nil {
		s.Log.Warnf("failed get data by id : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "data not found")
	}

	return converters.CategoryToResponse(category), nil
}

func (s CategoryServiceImpl) Create(request *dtos.CategoryCreateRequest) (*dtos.CategoryResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	contact := &entities.Category{
		Name: request.Name,
	}

	newContact, err := s.CategoryRepository.CreateData(contact)
	if err != nil {
		s.Log.Warnf("failed to create new data : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed create to create new data")
	}

	return converters.CategoryToResponse(newContact), nil
}

func (s CategoryServiceImpl) Update(request *dtos.CategoryUpdateRequest) (*dtos.CategoryResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	foundCategory, err := s.CategoryRepository.FindById(request.ID)
	if err != nil {
		s.Log.Warnf("failed get data by id : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "data not found")
	}

	foundCategory.Name = request.Name

	updatedCategory, err := s.CategoryRepository.UpdateData(foundCategory)
	if err != nil {
		s.Log.Warnf("failed to update data : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to update category")
	}

	return converters.CategoryToResponse(updatedCategory), nil
}

func (s CategoryServiceImpl) Delete(categoryId string) error {
	foundCategory, err := s.CategoryRepository.FindById(categoryId)
	if err != nil {
		s.Log.Warnf("failed get data by id : %+v", err)
		return fiber.NewError(fiber.StatusNotFound, "data not found")
	}

	if err := s.CategoryRepository.DeleteData(foundCategory); err != nil {
		s.Log.Warnf("failed to delete data : %+v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
