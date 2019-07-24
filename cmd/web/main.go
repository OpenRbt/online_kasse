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
	// Should start in API
	application.Start()

	// TO DO: transfer application to API
	server, err := api.NewWebServer()
	if err != nil {
		fmt.Println("Server start failure - program stopped")
		return
	}
	server.Start()
}
