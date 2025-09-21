package main

import (
	"github.com/BoomTHDev/tattoo_port/config"
	"github.com/BoomTHDev/tattoo_port/databases"
	"github.com/BoomTHDev/tattoo_port/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.ConnectionGetting().Begin()

	tattooMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return
	}
}

func tattooMigration(tx *gorm.DB) {
	tx.AutoMigrate(&entities.Tattoo{})
}
