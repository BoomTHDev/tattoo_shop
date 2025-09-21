package middleware

import (
	"errors"
	"log"

	"github.com/BoomTHDev/tattoo_port/pkg/custom"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if err == nil {
			return c.Next()
		}

		appErr := &custom.AppError{}
		fiberErr := &fiber.Error{}

		if errors.As(err, &appErr) {
			if appErr.StatusCode >= fiber.StatusInternalServerError && appErr.Err != nil {
				log.Printf("Internal AppError: %s, Original Error: %s\n", appErr.Message, appErr.Err.Error())
			} else if appErr.Err != nil {
				log.Printf("AppError: %s, Original Error: %s\n", appErr.Message, appErr.Err.Error())
			} else {
				log.Printf("AppError: %s\n", appErr.Message)
			}

			return c.Status(appErr.StatusCode).JSON(fiber.Map{
				"success": false,
				"message": appErr.Message,
			})
		}

		if errors.As(err, &fiberErr) {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": fiberErr.Message,
			})
		}

		log.Printf("Unhandled Error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "An unexpected internal server error occurred",
		})
	}
}
