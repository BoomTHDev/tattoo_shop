package http

import (
	_tattooController "github.com/BoomTHDev/tattoo_port/pkg/tattoo/controller"
	_tattooRepositoy "github.com/BoomTHDev/tattoo_port/pkg/tattoo/repository"
	_tattooService "github.com/BoomTHDev/tattoo_port/pkg/tattoo/service"
)

func (h *fiberServer) initTattooRouter() {
	// tattooRepository := _tattooRepositoy.NewTattooRepositoryImpl(h.db)
	tattooRepositoryMock := _tattooRepositoy.NewTattooRepositoryMockWithSeed()
	tattooService := _tattooService.NewTattooServiceImpl(tattooRepositoryMock)
	tattooController := _tattooController.NewTatttooController(tattooService)

	tattooRouter := h.app.Group("/v1/tattoos")
	tattooRouter.Get("/", tattooController.GetAllTattoos)
}
