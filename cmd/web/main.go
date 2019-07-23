package main

import (
	"fmt"
	"flag"
	"strconv"
	"reflect"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/powerman/structlog"
)

var log *structlog.Logger

func ProcessReceipt(ctx *fasthttp.RequestCtx) {
	price_interface := ctx.UserValue("sum")
	isCard_interface := ctx.UserValue("iscard")

	price_type := reflect.TypeOf(price_interface)
	if price_type.Name() != "string" {
		fmt.Fprintf(ctx, "Invaild type: got %s, but might be float64 in string", price_type)
		return
	}

	bankCard_type := reflect.TypeOf(isCard_interface)
	if bankCard_type.Name() != "string" {
		fmt.Fprintf(ctx, "Invalid type: got %s, but might be bool in string", bankCard_type)
		return
	}

	price_str := price_interface.(string)
	isBankCard_str := isCard_interface.(string)
	
	invalidType := false
	price, err := strconv.ParseFloat(price_str, 64)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of first parameter: might be float64\n")
		invalidType = true
	}
	isBankCard, err := strconv.ParseBool(isBankCard_str)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of second parameter: might be bool\n")
		invalidType = true
	}

	if invalidType {
		return
	}

	printer, err := app.NewWebApp()
	if err != nil {
		log.Fatalf("Error while initializing a cash control device %v", err)
	}

	printer.PrintReceipt(price, isBankCard)
	fmt.Println("Receipt with", price_str, "RUB and Bank card state:", isBankCard_str, "- sent to device")
}

func main() {
	flag.Parse()

	log = structlog.New()
	router := fasthttprouter.New()
	router.PUT("/:sum/:iscard", ProcessReceipt)

	port := ":8080"
	fmt.Println("Server is starting on port", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))	
}
