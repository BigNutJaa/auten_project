package main

import (
	"fmt"
	"github.com/BigNutJaa/users/internals/container"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	// add load .env for token
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// original main
	server, err := container.NewContainer()
	if err != nil {
		log.Panic(err)
	}

	if err := server.MigrateDB(); err != nil {
		log.Panic(err)
	}

	if err := server.Start(); err != nil {
		log.Panic(err)
	}
}
