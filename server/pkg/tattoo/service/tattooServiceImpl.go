package service

import (
	"github.com/BoomTHDev/tattoo_port/pkg/custom"
	_tattooModel "github.com/BoomTHDev/tattoo_port/pkg/tattoo/model"
	_tattooReposory "github.com/BoomTHDev/tattoo_port/pkg/tattoo/repository"
)

type tattooServiceImpl struct {
	tattooRepo _tattooReposory.TattooRepository
}

func NewTattooServiceImpl(tattooRepo _tattooReposory.TattooRepository) TattooService {
	return &tattooServiceImpl{
		tattooRepo: tattooRepo,
	}
}

func (s *tattooServiceImpl) GetAllTattoo() ([]_tattooModel.Tattoo, *custom.AppError) {
	tattoos, err := s.tattooRepo.GetAll()
	if err != nil {
		if custom.IsRecordFoundError(err) {
			return nil, custom.ErrNotFound("Tattoo not found.", err)
		}
		return nil, custom.ErrIntervalServer("Failed to get all tattoos.", err)
	}

	return _tattooModel.ToTattooModels(tattoos), nil
}
