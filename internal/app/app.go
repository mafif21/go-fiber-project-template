package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-fiber-project-template/internal/controllers"
	"go-fiber-project-template/internal/http/route"
	"go-fiber-project-template/internal/repositories"
	"go-fiber-project-template/internal/services"
	"gorm.io/gorm"
)

type StartAppConfig struct {
	App       *fiber.App
	DB        *gorm.DB
	Validator *validator.Validate
	Config    *viper.Viper
	Log       *logrus.Logger
}

func StartApp(config *StartAppConfig) {
	categoryRepository := repositories.NewCategoryRepositoryImpl(config.DB, config.Log)
	categoryservice := services.NewCategoryServiceImpl(config.Validator, categoryRepository, config.Log)
	categoryController := controllers.NewCategoryControllerImpl(categoryservice, config.Log)

	productRepository := repositories.NewProductRepositoryImpl(config.DB, config.Log)
	productService := services.NewProductServiceImpl(config.Validator, config.Log, productRepository)
	productController := controllers.NewProductControllerImpl(config.Log, productService)

	routeConfig := route.RouteConfig{
		App:                config.App,
		CategoryController: categoryController,
		ProductController:  productController,
	}

	routeConfig.Setup()
}
