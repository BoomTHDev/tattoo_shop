package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/BoomTHDev/tattoo_port/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	once                     sync.Once
	postgresDatabaseInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dns := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Bangkok",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.Port,
			conf.SSLMode,
		)

		conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s\n", conf.DBName)

		postgresDatabaseInstance = &postgresDatabase{conn}
	})

	return postgresDatabaseInstance
}

func (db *postgresDatabase) ConnectionGetting() *gorm.DB {
	return postgresDatabaseInstance.DB
}
