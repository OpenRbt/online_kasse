package app

import (
	"errors"
	"strconv"
	"sync"

	"github.com/DiaElectronics/online_kasse/cmd/web/dal"
	"github.com/DiaElectronics/online_kasse/cmd/web/fptr10"
)

// Data Access Layer is an interface for DAL usage
type DataAccessLayer interface {
	GetPrices(*GetData) (*ReceiptList, error)
	GetBankCards(*GetData) (*ReceiptList, error)
	GetCash(*GetData) (*ReceiptList, error)
	GetProcessed(*GetData) (*ReceiptList, error)
	GetUnprocessed(*GetData) (*ReceiptList, error)

	CreateReceipt(*Receipt) (*Receipt, error)

	DeleteReceipt(*Receipt) (*Receipt, error)
}

var (
	// ErrCannotConnect uses to describe Cash Register Device failures
	ErrCannotConnect = errors.New("Connection to Cash Register Device failed")
)

// Application is communicating with Cash Register Device
type Application struct {
	mutex sync.Mutex
	DB    DataAccessLayer
}

// RegisterReceipt sends Receipt to DAL (for instance, database)
func (app *Application) RegisterReceipt(currentData *Receipt) {
	app.DB.CreateReceipt(currentData)
}

// ResetShift sends signal to close current shift and open new one
func (app *Application) ResetShift() error {
	app.mutex.Lock()
	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()
	if !fptr.IsOpened() {
		return ErrCannotConnect
	}

	fptr.SetParam(1021, "Кассир Иванов И.")
	fptr.SetParam(1203, "123456789047")
	fptr.OperatorLogin()

	fptr.SetParam(fptr10.LIBFPTR_PARAM_REPORT_TYPE, fptr10.LIBFPTR_RT_CLOSE_SHIFT)
	fptr.Report()

	fptr.OpenShift()

	fptr.Close()

	fptr.Destroy()
	app.mutex.Unlock()

	return nil
}

// PingDevice checks connection to the Cash Register Device
func (app *Application) PingDevice() error {
	app.mutex.Lock()
	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()
	if !fptr.IsOpened() {
		fptr.Destroy()
		return ErrCannotConnect
	}
	fptr.Close()

	fptr.Destroy()

	app.mutex.Unlock()
	return nil
}

// PrintReceipt is sending receipt data to Cash Register Device
func (app *Application) PrintReceipt(price float64, isBankCard bool) error {

	app.mutex.Lock()
	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()
	if !fptr.IsOpened() {
		return ErrCannotConnect
	}

	fptr.SetParam(1021, "Кассир Иванов И.")
	fptr.SetParam(1203, "123456789047")
	fptr.OperatorLogin()

	fptr.OpenShift()

	fptr.SetParam(fptr10.LIBFPTR_PARAM_RECEIPT_TYPE, fptr10.LIBFPTR_RT_SELL)
	fptr.OpenReceipt()

	fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, "Мойка автомобиля")
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, price)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, fptr10.LIBFPTR_TAX_NO)
	fptr.Registration()

	if isBankCard {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_ELECTRONICALLY)
	} else {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_CASH)
	}

	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, price)
	fptr.Payment()

	fptr.CloseReceipt()
	fptr.CheckDocumentClosed()

	fptr.Close()

	fptr.Destroy()
	app.mutex.Unlock()

	return nil
}

// NewApplication constructs Application
func NewApplication(m sync.Mutex) (*Application, error) {
	res := &Application{}
	res.mutex = m
	res.DB = dal.NewPostgresDAL("anton", "", "localhost:5432")

	return res, nil
}

// Start initializes receipt processing goroutine
func (app *Application) Start() {

	// TO DO: start goroutine with data processing from DB
}
