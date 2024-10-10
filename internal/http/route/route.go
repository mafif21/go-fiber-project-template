package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-project-template/internal/controllers"
)

type RouteConfig struct {
	App                *fiber.App
	CategoryController controllers.CategoryController
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")
	c.RouteGuest(api)
	c.RouteAuth(api)
}

func (c *RouteConfig) RouteGuest(api fiber.Router) {
	api.Get("/category", c.CategoryController.Get)
	api.Get("/category/:categoryId", c.CategoryController.GetById)
	api.Post("/category", c.CategoryController.Create)
	api.Put("/category/:categoryId/edit", c.CategoryController.Update)
	api.Delete("/category/:categoryId/delete", c.CategoryController.Delete)
}

func (c *RouteConfig) RouteAuth(api fiber.Router) {}
