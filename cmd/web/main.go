package main

import (
	"fmt"

	"github.com/DiaElectronics/online_kasse/cmd/web/api"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
)

func main() {

	application, err := app.NewApplication()
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
