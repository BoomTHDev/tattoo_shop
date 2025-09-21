package main

import (
	"github.com/BoomTHDev/tattoo_port/config"
	"github.com/BoomTHDev/tattoo_port/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	_ = db
	// http := http.NewFiberServer(conf, db)

	// http.Start()
}
