package entities

import (
	"time"

	"github.com/google/uuid"
)

type Tattoo struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title     string    `gorm:"type:text;not null"`
	ImageURL  []string  `gorm:"type:text[];not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
