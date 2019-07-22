package main

import (
	//"github.com/gramework/gramework"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
)

func main() {
	//server := gramework.New()

	//server.GET("/", "hello, grameworld")

	printer, _ := app.NewWebApp()
	printer.PrintReceipt(50.0, true)

	//server.ListenAndServe()
}