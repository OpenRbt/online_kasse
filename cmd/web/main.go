package main

import (
	"github.com/gramework/gramework"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/powerman/structlog"
)

func main() {
	log := structlog.New()
	server := gramework.New()

	server.GET("/", "hello, grameworld")

	printer, err := app.NewWebApp()
	if err!= nil {
		log.Fatalf("error while initializing a printer %v", err)
	}
	printer.PrintReceipt(50.0, true)

	server.ListenAndServe()
}