package main

import (
	//"github.com/gramework/gramework"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
)

func main() {
	//server := gramework.New()

	//server.GET("/", "hello, grameworld")

	printer := app.NewWebApp()
	printer.printReceipt(50.0, true)

	//server.ListenAndServe()
}