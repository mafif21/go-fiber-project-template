package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-project-template/internal/controllers"
)

type RouteConfig struct {
	App                *fiber.App
	CategoryController controllers.CategoryController
	ProductController  controllers.ProductController
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")
	c.RouteGuest(api)
	c.RouteAuth(api)
}

func (c *RouteConfig) RouteGuest(api fiber.Router) {
	api.Get("/category", c.CategoryController.GetCategory)
	api.Get("/category/:categoryId", c.CategoryController.GetCategoryById)
	api.Post("/category", c.CategoryController.CreateCategory)
	api.Put("/category/:categoryId/edit", c.CategoryController.UpdateCategory)
	api.Delete("/category/:categoryId/delete", c.CategoryController.DeleteCategory)

	api.Get("/product", c.ProductController.GetProduct)
	api.Get("/product/:productId", c.ProductController.GetProductById)
	api.Post("/product", c.ProductController.CreateProduct)
	api.Put("/product/:productId/edit", c.ProductController.UpdateProduct)
	api.Delete("/product/:productId/delete", c.ProductController.DeleteProduct)
}

func (c *RouteConfig) RouteAuth(api fiber.Router) {}
