package api

import (
	"fmt"
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

	server.application.RegisterReceipt(currentReceipt)
}

// Start initializes Web Server, starts application and begins serving
func (server *WebServer) Start() {
	server.log = structlog.New()

	server.application.Start()

	router := fasthttprouter.New()
	router.PUT("/:post/:sum/:iscard", server.PushReceipt)

	port := ":8080"

	fmt.Println("Server is starting on port", port)
	server.log.Fatal(fasthttp.ListenAndServeTLS(port, "cert.pem", "key.pem", router.Handler))
}

// NewWebServer constructs Web Server
func NewWebServer(application app.IncomeRegistration) (*WebServer, error) {
	res := &WebServer{}
	res.application = application

	return res, nil
}
