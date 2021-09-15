package bootstrap

import (
	"github.com/macmagic/technical-test-deporvillage/internal/application"
	"github.com/macmagic/technical-test-deporvillage/internal/infrastructure"
	"log"
)

const maxLifeTime = 30

func Run() {
	config, err := application.NewConfig("./app.json")

	if err != nil {
		log.Fatalln("Cannot load application config")
	}

	server := infrastructure.NewServer(config)
	server.StartListen()

	//time.Sleep(900 * time.Second)

	log.Println("Finish")
}
