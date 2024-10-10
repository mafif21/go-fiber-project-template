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
	c.RouteGuest()
	c.RouteAuth()
}

func (c *RouteConfig) RouteGuest() {
	c.App.Get("/api/category", c.CategoryController.Get)
	c.App.Get("/api/category/:categoryId", c.CategoryController.GetById)
	c.App.Post("/api/category", c.CategoryController.Create)
	c.App.Put("/api/category/:categoryId/edit", c.CategoryController.Update)
	c.App.Delete("/api/category/:categoryId/delete", c.CategoryController.Delete)
}

func (c *RouteConfig) RouteAuth() {}
