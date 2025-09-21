package controller

import (
	_tattooService "github.com/BoomTHDev/tattoo_port/pkg/tattoo/service"
	"github.com/gofiber/fiber/v2"
)

type tattooController struct {
	tattooService _tattooService.TattooService
}

func NewTatttooController(tattooService _tattooService.TattooService) *tattooController {
	return &tattooController{
		tattooService: tattooService,
	}
}

func (c *tattooController) GetAllTattoos(ctx *fiber.Ctx) error {
	tattoos, appErr := c.tattooService.GetAllTattoo()
	if appErr != nil {
		return appErr
	}

	return ctx.JSON(tattoos)
}
