package api

import (
	"fmt"
	"strconv"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/buaazp/fasthttprouter"
	"github.com/powerman/structlog"
	"github.com/valyala/fasthttp"
)

var log = structlog.New()

// WebServer accepts PUT requests with payload of Receipts and pushes them to Application
type WebServer struct {
	application app.IncomeRegistration
}

// Ping answers on any valid GET request as OK (code 200)
func (server *WebServer) Ping(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// PushReceipt pushes new Receipt to Application
func (server *WebServer) PushReceipt(ctx *fasthttp.RequestCtx) {
	currentReceipt := app.NewReceipt()

	postStr := ctx.UserValue("post").(string)
	priceStr := ctx.UserValue("sum").(string)
	isBankCardStr := ctx.UserValue("iscard").(string)

	invalidType := false

	post, err := strconv.ParseInt(postStr, 10, 64)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of post number: might be int64\n")
		invalidType = true
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of money amount: might be float64\n")
		invalidType = true
	}
	isBankCard, err := strconv.ParseBool(isBankCardStr)
	if err != nil {
		fmt.Fprintf(ctx, "Invalid type of bank card flag: might be bool (or 0/1)\n")
		invalidType = true
	}

	if invalidType {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	currentReceipt.Post = post
	currentReceipt.Price = price
	currentReceipt.IsBankCard = isBankCard

	log.Info("API got new receipt")

	server.application.RegisterReceipt(currentReceipt)

	ctx.SetStatusCode(fasthttp.StatusOK)
}

// Start initializes Web Server, starts application and begins serving
func (server *WebServer) Start(errc chan<- error) {
	server.application.Start()

	router := fasthttprouter.New()
	router.PUT("/:post/:sum/:iscard", server.PushReceipt)
	router.GET("/ping_kasse", server.Ping)

	port := ":443"

	log.Info("Server is starting on port", port)
	errc <- fasthttp.ListenAndServeTLS(port, "cert.pem", "key.pem", router.Handler)
}

// NewWebServer constructs Web Server
func NewWebServer(application app.IncomeRegistration) *WebServer {
	res := &WebServer{}
	res.application = application

	return res
}
