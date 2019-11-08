package main

import (
	"fmt"
	"sync"

	"github.com/DiaElectronics/online_kasse/cmd/web/api"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/DiaElectronics/online_kasse/cmd/web/dal"
	"github.com/DiaElectronics/online_kasse/cmd/web/device"
	"github.com/powerman/structlog"
)

var log = structlog.New()

func run(errc chan<- error) {
	var mutex sync.Mutex
	db, err := dal.NewPostgresDAL("kaznachey", "RfpyfxtqAF", "localhost:5432")
	if err != nil {
		errc <- err
		return
	}

	dev, err := device.NewKaznacheyFA(mutex)
	if err != nil {
		errc <- err
		return
	}

	application := app.NewApplication(db, dev, errc)
	server := api.NewWebServer(application)

	server.Start(errc)
}

func main() {
	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(structlog.DBG))
	log.Info("Server is preparing to start...")

	errc := make(chan error)
	go run(errc)
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

}
