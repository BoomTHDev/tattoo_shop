package model

import (
	"github.com/BoomTHDev/tattoo_port/entities"
	"github.com/google/uuid"
)

type (
	Tattoo struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		ImageURL []string  `json:"image_url"`
	}

	CreateTattooReq struct {
		Title    string   `json:"title"`
		ImageURL []string `json:"image_url"`
	}
)

func ToTattooModel(tattoo *entities.Tattoo) *Tattoo {
	if tattoo == nil {
		return nil
	}

	return &Tattoo{
		ID:       tattoo.ID,
		Title:    tattoo.Title,
		ImageURL: tattoo.ImageURL,
	}
}

func ToTattooModels(tattoos []entities.Tattoo) []Tattoo {
	if tattoos == nil {
		return nil
	}

	tattooModels := []Tattoo{}
	for _, tattoo := range tattoos {
		tattooModels = append(tattooModels, Tattoo{
			ID:       tattoo.ID,
			Title:    tattoo.Title,
			ImageURL: tattoo.ImageURL,
		})
	}

	return tattooModels
}
