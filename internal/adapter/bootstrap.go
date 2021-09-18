package adapter

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
	//Load app config
	appConfig, err := config.NewConfig("./app.json")

	if err != nil {
		log.Fatalln("Cannot load application configuration", err.Error())
	}

	// Dependencies Injection
	repository := infrastructure.NewFileRepository(appConfig)
	skuService := domain.NewSkuService(repository)
	server := infrastructure.NewServer(appConfig, skuService)

	//Stop channel to stop the main routine when receives a "terminate" command or the lifetime terminate
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
		log.Println(skuService.GenerateReport())
		repository.CloseFile()
		os.Exit(0)
	}
}

func runnerControl(stopApp chan bool) {
	time.Sleep(maxLifeTime * time.Second)
	stopApp <- true
}