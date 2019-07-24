package api

import "github.com/valyala/fasthttp"
import "github.com/buaazp/fasthttprouter"

var log *structlog.Logger

type WebServer struct {

}

type Receipt struct {
	Price float64
	IsBankCard bool
}

func NewReceipt () (*Receipt, error) {
	res := &Receipt{}
	return res, nil	
}

func (server *WebServer) SubmitReceipt(ctx *fasthttp.RequestCtx) {
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

	printer, err := app.NewWebApp()
	if err != nil {
		log.Fatalf("Error while initializing a cash control device %v", err)
	}

	printer.PrintReceipt(price, isBankCard)
	fmt.Println("Receipt with", price_str, "RUB and Bank card state:", isBankCard_str, "- sent to device")
}

func NewWebServer () (*WebServer, error) {
	res := &WebServer{}

	router := fasthttprouter.New()
	router.PUT("/:sum/:iscard", ProcessReceipt)

	port := ":8080"

	fmt.Println("Server is starting on port", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))

	return res, nil	
}