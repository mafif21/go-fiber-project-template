package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/config"
	"go-fiber-project-template/internal/model/dtos"
	"go-fiber-project-template/internal/services"
)

type ProductController interface {
	GetProduct(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	Log            *logrus.Logger
	ProductService services.ProductService
}

func NewProductControllerImpl(log *logrus.Logger, productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		Log:            log,
		ProductService: productService,
	}
}

func (c ProductControllerImpl) GetProduct(ctx *fiber.Ctx) error {
	products, err := c.ProductService.Get()
	if err != nil {
		c.Log.Warnf("failed to get all products : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[[]dtos.ProductResponse]{
		Status:  fiber.StatusOK,
		Message: "success get all categories",
		Data:    products,
	})
}

func (c ProductControllerImpl) GetProductById(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	product, err := c.ProductService.GetById(productId)
	if err != nil {
		c.Log.Warnf("failed to get detail product : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.ProductResponse]{
		Status:  fiber.StatusOK,
		Message: fmt.Sprintf("success get product with id %s", product.ID),
		Data:    product,
	})
}

func (c ProductControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	request := new(dtos.ProductCreateRequest)

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	newProduct, err := c.ProductService.Create(request)
	if err != nil {
		c.Log.Warnf("failed to create new product : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.ProductResponse]{
		Status:  fiber.StatusOK,
		Message: "success create new product",
		Data:    newProduct,
	})
}

func (c ProductControllerImpl) UpdateProduct(ctx *fiber.Ctx) error {
	request := new(dtos.ProductUpdateRequest)

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	productId := ctx.Params("productId")
	request.ID = productId

	updatedProduct, err := c.ProductService.Update(request)
	if err != nil {
		c.Log.Warnf("failed to update existing product : %+v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.ProductResponse]{
		Status:  fiber.StatusOK,
		Message: "success edit existing product",
		Data:    updatedProduct,
	})
}

func (c ProductControllerImpl) DeleteProduct(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	err := c.ProductService.Delete(productId)
	if err != nil {
		c.Log.Warnf("failed to delete existing product : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.ProductResponse]{
		Status:  fiber.StatusOK,
		Message: "success delete existing data",
	})
}
