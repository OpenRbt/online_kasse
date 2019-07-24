package main

import (
	"fmt"
	"flag"
	"github.com/DiaElectronics/online_kasse/cmd/web/api"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
)

func main() {
	flag.Parse()

	application, err := NewApplication()
	if err != nil {
		fmt.Println("Application start failure - program stopped")
		return
	}
	application.Start()	

	// TO DO: transfer application to API
	server, err := NewWebServer()
	if err != nil {
		fmt.Println("Server start failure - program stopped")
		return
	}
	server.Start()
}
