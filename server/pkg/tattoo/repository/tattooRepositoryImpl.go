package repository

import (
	"github.com/BoomTHDev/tattoo_port/databases"
	"github.com/BoomTHDev/tattoo_port/entities"
)

type tattooRepositoryImpl struct {
	db databases.Database
}

func NewTattooRepositoryImpl(db databases.Database) TattooRepository {
	return &tattooRepositoryImpl{db}
}

func (r *tattooRepositoryImpl) Create(tattoo *entities.Tattoo) error {
	conn := r.db.ConnectionGetting()

	if err := conn.Create(&tattoo).Error; err != nil {
		return err
	}

	return nil
}

func (r *tattooRepositoryImpl) GetAll() ([]entities.Tattoo, error) {
	conn := r.db.ConnectionGetting()

	tattoos := []entities.Tattoo{}

	if err := conn.Find(&tattoos).Error; err != nil {
		return nil, err
	}

	return tattoos, nil
}

func (r *tattooRepositoryImpl) GetById(id string) (*entities.Tattoo, error) {
	conn := r.db.ConnectionGetting()

	tattoo := entities.Tattoo{}

	if err := conn.First(&tattoo, id).Error; err != nil {
		return nil, err
	}

	return &tattoo, nil
}
