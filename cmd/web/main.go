package main

import (
	"fmt"
	"sync"

	"github.com/DiaElectronics/online_kasse/cmd/web/api"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
)

func main() {
	var mutex sync.Mutex

	application, err := app.NewApplication(mutex)
	if err != nil {
		fmt.Println("Application start failure - program stopped")
		return
	}

	server, err := api.NewWebServer(application)
	if err != nil {
		fmt.Println("Server start failure - program stopped")
		return
	}
	server.Start()
}
