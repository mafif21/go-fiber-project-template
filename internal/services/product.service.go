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

type ProductService interface {
	Get() ([]dtos.ProductResponse, error)
	GetById(categoryId string) (*dtos.ProductResponse, error)
	Create(request *dtos.ProductCreateRequest) (*dtos.ProductResponse, error)
	Update(request *dtos.ProductUpdateRequest) (*dtos.ProductResponse, error)
	Delete(categoryId string) error
}

type ProductServiceImpl struct {
	Validate          *validator.Validate
	Log               *logrus.Logger
	ProductRepository repositories.ProductRepository
}

func NewProductServiceImpl(validate *validator.Validate, log *logrus.Logger, productRepository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		Validate:          validate,
		Log:               log,
		ProductRepository: productRepository,
	}
}

func (s ProductServiceImpl) Get() ([]dtos.ProductResponse, error) {
	products, err := s.ProductRepository.FindAll()
	if err != nil {
		s.Log.Warnf("failed get all data categories : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "invalid database")
	}

	response := make([]dtos.ProductResponse, len(products))
	for i, product := range products {
		response[i] = *converters.ProductToResponse(&product)
	}
	return response, nil
}

func (s ProductServiceImpl) GetById(categoryId string) (*dtos.ProductResponse, error) {
	product, err := s.ProductRepository.FindById(categoryId)
	if err != nil {
		s.Log.Warnf("failed get data by id : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "data not found")
	}

	return converters.ProductToResponse(product), nil
}

func (s ProductServiceImpl) Create(request *dtos.ProductCreateRequest) (*dtos.ProductResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	product := &entities.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryID: request.CategoryID,
	}

	newProduct, err := s.ProductRepository.CreateData(product)
	if err != nil {
		s.Log.Warnf("failed to create new data : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed create to create new data")
	}

	return converters.ProductToResponse(newProduct), nil
}

func (s ProductServiceImpl) Update(request *dtos.ProductUpdateRequest) (*dtos.ProductResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	foundProduct, err := s.ProductRepository.FindById(request.ID)
	if err != nil {
		s.Log.Warnf("failed get data by id : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "data not found")
	}

	foundProduct.Name = request.Name
	foundProduct.Price = request.Price
	foundProduct.CategoryID = request.CategoryID

	updatedProduct, err := s.ProductRepository.UpdateData(foundProduct)
	if err != nil {
		s.Log.Warnf("failed to update data : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "failed to update category")
	}

	return converters.ProductToResponse(updatedProduct), nil
}

func (s ProductServiceImpl) Delete(categoryId string) error {
	product, err := s.ProductRepository.FindById(categoryId)
	if err != nil {
		s.Log.Warnf("failed get data by id : %+v", err)
		return fiber.NewError(fiber.StatusNotFound, "data not found")
	}

	if err := s.ProductRepository.DeleteData(product); err != nil {
		s.Log.Warnf("failed to delete data : %+v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
