package service

import (
	"github.com/BoomTHDev/tattoo_port/pkg/custom"
	_tattooModel "github.com/BoomTHDev/tattoo_port/pkg/tattoo/model"
)

type TattooService interface {
	// NewTattoo(tattoo *_tattooModel.CreateTattooReq) (*_tattooModel.Tattoo, *custom.AppError)
	GetAllTattoo() ([]_tattooModel.Tattoo, *custom.AppError)
	// GetTattooById(id string) (*_tattooModel.Tattoo, *custom.AppError)
}
