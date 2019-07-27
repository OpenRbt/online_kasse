package app

import (
	"errors"
	"fmt"
)

// IncomeRegistration is an interface for accepting income Receipts from Web Server
type IncomeRegistration interface {
	RegisterReceipt(*Receipt)
	Start()
}

// DataAccessLayer is an interface for DAL usage from Application
type DataAccessLayer interface {
	GetByPrice(QueryData) (*ReceiptList, error)
	GetWithBankCards(QueryData) (*ReceiptList, error)
	GetWithCash(QueryData) (*ReceiptList, error)
	GetUnprocessedOnly(QueryData) (*ReceiptList, error)
	GetProcessedOnly(QueryData) (*ReceiptList, error)

	Create(*Receipt) (*Receipt, error)

	UpdateStatus(*Receipt) (bool, error)

	DeleteByID(int64) (int64, error)
}

// DeviceAccessLayer is an interface for DevAL usage from Application
type DeviceAccessLayer interface {
	ResetShift() error
	PrintReceipt(*Receipt) error
	PingDevice() error
}

// ErrCannotConnect describes DevAL failures
var ErrCannotConnect = errors.New("Device Access Layer is unavailable")

// Application is responsible for all logics and communicates with other layers
type Application struct {
	DB     DataAccessLayer
	Device DeviceAccessLayer
}

// RegisterReceipt sends Receipt to DAL for saving/registration
func (app *Application) RegisterReceipt(currentData *Receipt) {
	app.DB.Create(currentData)
	list, _ := app.DB.GetUnprocessedOnly(QueryData{Limit: 100, LastId: 0})
	fmt.Println(list)
}

// NewApplication constructs Application
func NewApplication(db DataAccessLayer, dev DeviceAccessLayer) (*Application, error) {
	res := &Application{}

	res.DB = db
	res.Device = dev

	return res, nil
}

// Start initializes Receipt Processing goroutine
func (app *Application) Start() {

	// TO DO: start goroutine with data processing from DB
}
