package repository

import "github.com/BoomTHDev/tattoo_port/entities"

type TattooRepository interface {
	Create(tattoo *entities.Tattoo) error
	GetAll() ([]entities.Tattoo, error)
	GetById(id string) (*entities.Tattoo, error)
}
