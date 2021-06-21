package app

import (
	"errors"
	"time"

	"github.com/powerman/structlog"
)

var log = structlog.New()

// IncomeRegistration is an interface for accepting income Receipts from Web Server
type IncomeRegistration interface {
	RegisterReceipt(*Receipt)
	Start()
	Info() string
}

// DataAccessLayer is an interface for DAL usage from Application
type DataAccessLayer interface {
	GetByPost(QueryData) (*ReceiptList, error)
	GetWithBankCards(QueryData) (*ReceiptList, error)
	GetWithCash(QueryData) (*ReceiptList, error)
	GetUnprocessedOnly(QueryData) (*ReceiptList, error)
	GetProcessedOnly(QueryData) (*ReceiptList, error)

	Create(*Receipt) (*Receipt, error)

	UpdateStatus(Receipt) (bool, error)

	DeleteByID(int64) (int64, error)
	Info() string
}

// DeviceAccessLayer is an interface for DevAL usage from Application
type DeviceAccessLayer interface {
	PrintReceipt(Receipt) error
	PingDevice() error
}

// Errors for DevAL failures
var (
	ErrCannotConnect              = errors.New("Cash Register Device is unable to connect")
	ErrSetupFailure               = errors.New("Connection setup failed")
	ErrLoginFailure               = errors.New("Operator login failed")
	ErrShiftCloseFailure          = errors.New("Shift closing failed")
	ErrShiftOpenFailure           = errors.New("Shift opening failed")
	ErrShiftState                 = errors.New("Get shift state")
	ErrReceiptCreationFailure     = errors.New("Receipt opening failed")
	ErrReceiptRegistrationFailure = errors.New("Receipt registration failed")
	ErrTotalRegistrationFailure   = errors.New("Total registration failed")
	ErrPaymentSetFailure          = errors.New("Payment method set failed")
	ErrReceiptCloseFailure        = errors.New("Receipt close failed")
	ErrUnableToGetFiscalData      = errors.New("Unable to get fiscal data")
	ErrCannotDisconnect           = errors.New("Cash Register Device in unable to disconnect")
	ErrNotFound                   = errors.New("not found")
	ErrDeviceReboot               = errors.New("device reboot")
)

// Application is responsible for all logics and communicates with other layers
type Application struct {
	DB     DataAccessLayer
	Device DeviceAccessLayer
	errc   chan<- error
}

// RegisterReceipt sends Receipt to DAL for saving/registration
func (app *Application) RegisterReceipt(currentData *Receipt) {
	_, err := app.DB.Create(currentData)

	if err != nil {
		app.errc <- err
		return
	}
	log.Info("New receipt added to DB")
}

// Info returns database information
func (app *Application) Info() string {
	return app.DB.Info()
}

// NewApplication constructs Application
func NewApplication(db DataAccessLayer, dev DeviceAccessLayer, errchannel chan<- error) *Application {
	res := &Application{}

	res.DB = db
	res.Device = dev
	res.errc = errchannel

	return res
}

func (app *Application) loop() {
	needToSleep := false
	for {
		listToProcess, err := app.DB.GetUnprocessedOnly(QueryData{Limit: 1, LastID: 0})
		if listToProcess == nil || err != nil || app.Device == nil {
			log.Info("List of unprocessed receipts is empty")
			time.Sleep(time.Second * 5)
			continue
		}

		if listToProcess.Total != 0 {
			needToSleep = false

			receiptToProcess := listToProcess.Receipts[0]
			err := app.Device.PrintReceipt(receiptToProcess)
			if err != nil {
				log.Info("Error while printing a receipt")
				time.Sleep(time.Second * 5)
				continue
			}
			log.Info("Receipt printed")
			if app.DB.Info() == "memdb" {
				_, err = app.DB.DeleteByID(receiptToProcess.ID)
				if err != nil {
					log.Info("Error while deleting a receipt")
					continue
				}
				log.Info("Receipt deleted in DB")
			} else {
				_, err = app.DB.UpdateStatus(receiptToProcess)
				if err != nil {
					log.Info("Error while updating a receipt")
					continue
				}
				log.Info("Receipt updated in DB")
			}
		} else {
			needToSleep = true
		}

		if needToSleep {
			time.Sleep(time.Second * 5)
		}
	}
}

// Start initializes Receipt Processing goroutine
func (app *Application) Start() {
	go app.loop()
}
