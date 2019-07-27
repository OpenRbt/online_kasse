package main

import (
	"fmt"
	"sync"

	"github.com/DiaElectronics/online_kasse/cmd/web/api"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/DiaElectronics/online_kasse/cmd/web/dal"
	"github.com/DiaElectronics/online_kasse/cmd/web/device"
)

func main() {
	var mutex sync.Mutex
	db, err := dal.NewPostgresDAL("kaznachey", "test", "127.0.0.1:5432")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Data base start failure - program stopped")
		return
	}

	dev, err := device.NewKaznacheyFA(mutex)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Device start failure - program stopped")
		return
	}

	application, err := app.NewApplication(db, dev)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Application start failure - program stopped")
		return
	}

	server, err := api.NewWebServer(application)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Server start failure - program stopped")
		return
	}
	server.Start()
}
