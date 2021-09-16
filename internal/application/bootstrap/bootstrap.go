package bootstrap

import (
	"github.com/macmagic/technical-test-deporvillage/internal/application/config"
	"github.com/macmagic/technical-test-deporvillage/internal/domain"
	"github.com/macmagic/technical-test-deporvillage/internal/infrastructure"
	"log"
	"os"
	"time"
)

const maxLifeTime = 900

func Run() {

	appConfig, err := config.NewConfig("./app.json")

	if err != nil {
		log.Fatalln("Cannot load application appConfig")
	}

	repository := infrastructure.NewFileRepository()
	skuService := domain.NewSkuService(repository)

	server := infrastructure.NewServer(appConfig, skuService)

	stopApp := make(chan bool)
	go runnerControl(stopApp)

	go func() {
		server.StartListen()
		stopApp <- true
	}()

	select {
	case <-stopApp:
		server.StopServer()
		log.Println("Stopping application...")
		os.Exit(0)
	}
}

func runnerControl(stopApp chan bool) {
	time.Sleep(maxLifeTime * time.Second)
	stopApp <- true
}
