package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DiaElectronics/online_kasse/cmd/web/models"
	"github.com/buaazp/fasthttprouter"
	"github.com/powerman/structlog"
	"github.com/valyala/fasthttp"
)

// WebServer maintains PUT requestes and sends data to Application
type WebServer struct {
	log *structlog.Logger
}

// PushReceipt pushes new receipt object to Application
func (server *WebServer) PushReceipt(ctx *fasthttp.RequestCtx) {
	currentReceipt, err := models.NewReceipt()
	if err != nil {
		log.Fatalf("Error while creating a new receipt")
	}

	priceStr := ctx.UserValue("sum").(string)
	isBankCardStr := ctx.UserValue("iscard").(string)

	invalidType := false
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of first parameter: might be float64\n")
		invalidType = true
	}
	isBankCard, err := strconv.ParseBool(isBankCardStr)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of second parameter: might be bool\n")
		invalidType = true
	}

	if invalidType {
		return
	}

	currentReceipt.Price = price
	currentReceipt.IsBankCard = isBankCard

	// TO DO: send receipt to APPLICATION

	/*
		printer, err := app.KaznacheyFA()
		if err != nil {
			log.Fatalf("Error while initializing a cash control device %v", err)
		}

		printer.PrintReceipt(price, isBankCard)
		fmt.Println("Receipt with", price_str, "RUB and Bank card state:", isBankCard_str, "- sent to device")
	*/
}

// Start initializes WebServer
func (server *WebServer) Start() {
	server.log = structlog.New()

	// start application here
	router := fasthttprouter.New()
	router.PUT("/:sum/:iscard", server.PushReceipt)

	port := ":8080"

	fmt.Println("Server is starting on port", port)
	server.log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}

// NewWebServer constructs WebServer object
func NewWebServer() (*WebServer, error) {
	res := &WebServer{}

	return res, nil
}
