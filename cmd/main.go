package main

import (
	"github.com/macmagic/technical-test-deporvillage/internal/adapter"
	"github.com/macmagic/technical-test-deporvillage/internal/application/config"
	"log"
)

func main() {

	log.Println("Starting application!")

	//Load app config
	appConfig, err := config.NewConfig("./app.json")

	if err != nil {
		log.Fatalln("Cannot load application configuration", err.Error())
	}

	adapter.Run(appConfig)
}
