package api

import (
	"fmt"
	"log"
	"strconv"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/buaazp/fasthttprouter"
	"github.com/powerman/structlog"
	"github.com/valyala/fasthttp"
)

// WebServer accepts PUT requests with payload of Receipts and pushes them to Application
type WebServer struct {
	log         *structlog.Logger
	application app.IncomeRegistration
}

// PushReceipt pushes new Receipt to Application
func (server *WebServer) PushReceipt(ctx *fasthttp.RequestCtx) {
	currentReceipt, err := app.NewReceipt()
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

	server.application.RegisterReceipt(currentReceipt)
}

// Start initializes Web Server, starts application and begins serving
func (server *WebServer) Start() {
	server.log = structlog.New()

	server.application.Start()

	router := fasthttprouter.New()
	router.PUT("/:sum/:iscard", server.PushReceipt)

	port := ":8080"

	fmt.Println("Server is starting on port", port)
	server.log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}

// NewWebServer constructs Web Server
func NewWebServer(application app.IncomeRegistration) (*WebServer, error) {
	res := &WebServer{}
	res.application = application

	return res, nil
}
