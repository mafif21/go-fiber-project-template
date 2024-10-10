package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-fiber-project-template/internal/model/dtos"
)

func ErrorHandler(ctx *fiber.Ctx, log *logrus.Logger, err error) error {
	if err, ok := err.(*fiber.Error); ok {
		return ctx.Status(err.Code).JSON(&dtos.ErrorResponse{
			Status:  err.Code,
			Message: err.Message,
		})
	}

	log.Warnf("unexpected error : %v", err)
	return ctx.Status(fiber.StatusInternalServerError).JSON(&dtos.ErrorResponse{
		Status:  fiber.StatusInternalServerError,
		Message: "unexpected error ",
	})
}
