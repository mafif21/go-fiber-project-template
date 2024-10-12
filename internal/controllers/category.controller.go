package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/config"
	"go-fiber-project-template/internal/model/dtos"
	"go-fiber-project-template/internal/services"
	"math"
)

type CategoryController interface {
	GetCategory(ctx *fiber.Ctx) error
	GetCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
	Log             *logrus.Logger
}

func NewCategoryControllerImpl(categoryService services.CategoryService, log *logrus.Logger) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService, Log: log}
}

func (c CategoryControllerImpl) GetCategory(ctx *fiber.Ctx) error {
	request := &dtos.CategorySearchRequest{
		Name: ctx.Query("name", ""),
		SearchConf: &dtos.SearchConf{
			Page: ctx.QueryInt("page", 1),
			Size: ctx.QueryInt("size", 10),
		},
	}

	categories, total, err := c.CategoryService.Get(request)
	if err != nil {
		c.Log.Warnf("failed to get all categoies : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	pagination := &dtos.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(&dtos.WebResponse[[]dtos.CategoryResponse]{
		Status:  fiber.StatusOK,
		Message: "success get all categories",
		Data:    categories,
		Paging:  pagination,
	})
}

func (c CategoryControllerImpl) GetCategoryById(ctx *fiber.Ctx) error {
	category, err := c.CategoryService.GetById(ctx.Params("categoryId"))

	if err != nil {
		c.Log.Warnf("failed to get detail category : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.CategoryResponse]{
		Status:  fiber.StatusOK,
		Message: fmt.Sprintf("success get category with id %s", category.ID),
		Data:    category,
	})
}

func (c CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	request := new(dtos.CategoryCreateRequest)

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	category, err := c.CategoryService.Create(request)
	if err != nil {
		c.Log.Warnf("failed to create new category : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.CategoryResponse]{
		Status:  fiber.StatusCreated,
		Message: "success create new data",
		Data:    category,
	})
}

func (c CategoryControllerImpl) UpdateCategory(ctx *fiber.Ctx) error {
	request := new(dtos.CategoryUpdateRequest)

	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("failed to parse request body : %+v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	request.ID = ctx.Params("categoryId")

	category, err := c.CategoryService.Update(request)
	if err != nil {
		c.Log.Warnf("failed to update existing category : %+v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.CategoryResponse]{
		Status:  fiber.StatusOK,
		Message: "success update existing data",
		Data:    category,
	})
}

func (c CategoryControllerImpl) DeleteCategory(ctx *fiber.Ctx) error {
	err := c.CategoryService.Delete(ctx.Params("categoryId"))
	if err != nil {
		c.Log.Warnf("failed to delete existing category : %v", err)
		return config.ErrorHandler(ctx, c.Log, err)
	}

	return ctx.JSON(&dtos.WebResponse[*dtos.CategoryResponse]{
		Status:  fiber.StatusOK,
		Message: "success delete existing data",
	})
}
