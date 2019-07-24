package api

import (
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/powerman/structlog"
)

type WebServer struct {
	// send to constructor, change Log to log
	log *structlog.Logger
}

func (server *WebServer) PushReceipt(ctx *fasthttp.RequestCtx) {
	currentReceipt, err := NewReceipt()
	if err != nil {
		log.Fatalf("Error while creating a new receipt")
	}

	price_str := ctx.UserValue("sum").(string)
	isBankCard_str := ctx.UserValue("iscard").(string)
	
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

func (server *WebServer) Start() {
	server.Log = structlog.New()

	// start application here
	router := fasthttprouter.New()
	router.PUT("/:sum/:iscard", ProcessReceipt)

	port := ":8080"

	fmt.Println("Server is starting on port", port)
	server.Log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}

func NewWebServer () (*WebServer, error) {
	res := &WebServer{}
	
	return res, nil	
}