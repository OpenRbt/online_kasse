package device

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/DiaElectronics/online_kasse/cmd/web/fptr10"
	"github.com/powerman/structlog"
)

var log = structlog.New()

// Config kasse settings
type Config struct {
	Cashier         string
	CashierINN      string
	ReceiptItemName string
	Tax             int
}

// ConfigSvc is an interface for getting kasse settings
type ConfigSvc interface {
	GetConfig() (*Config, error)
}

// KaznacheyFA representes object of Device, connected by USB
type KaznacheyFA struct {
	mutex *sync.Mutex
	cfg   *Config
}

// PingDevice checks connection to the Device
//nolint
func (dev *KaznacheyFA) PingDevice() error {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	fptr := fptr10.New()
	if fptr == nil {
		return app.ErrCannotConnect
	}
	log.Info("device is initialized")
	defer fptr.Destroy()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_KAZNACHEY_FA))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))

	if err := fptr.ApplySingleSettings(); err != nil {
		log.Info(err)
		return app.ErrSetupFailure
	}

	if err := fptr.Open(); err != nil {
		log.Info(err)
		return app.ErrCannotConnect
	}

	fptr.Close()

	return nil
}

// PrintReceipt sends Receipt to the Device driver
//nolint
func (dev *KaznacheyFA) PrintReceipt(data app.Receipt) error {
	if dev == nil {
		fmt.Println("can't print on nil device")
		return app.ErrCannotConnect
	}
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	fptr := fptr10.New()
	defer fptr.Destroy()

	// Stage 1: Configure connection to Device
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_KAZNACHEY_FA))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))

	if err := fptr.ApplySingleSettings(); err != nil {
		log.Info(err)
		return app.ErrSetupFailure
	}

	// Stage 2: Connect to Device
	if err := fptr.Open(); err != nil {
		log.Info(err)
		return app.ErrCannotConnect
	}
	defer fptr.Close()

	log.Info("Connection to Cash Register Device opened")

	// Stage 3: Register the responsible person
	fptr.SetParam(1021, dev.cfg.Cashier)
	fptr.SetParam(1203, dev.cfg.CashierINN)
	if err := fptr.OperatorLogin(); err != nil {
		log.Info(err)
		return app.ErrLoginFailure
	}

	// Stage 4: Check the shift: open or close it (and open again)
	// If the shift was already opened - just do nothing
	fptr.OpenShift()
	errorCode := fptr.ErrorCode()

	// If shift expired (was more than 24 hours long) - close it and open again
	if errorCode == 68 || errorCode == 141 {
		log.Info("Shift expired - closing and reopening")

		fptr.SetParam(fptr10.LIBFPTR_PARAM_REPORT_TYPE, fptr10.LIBFPTR_RT_CLOSE_SHIFT)
		if err := fptr.Report(); err != nil {
			log.Info(err)
			return app.ErrShiftCloseFailure
		}

		fptr.OpenShift()
		errorCode := fptr.ErrorCode()
		if errorCode != 0 {
			log.Info("Error while opening shift")
			return app.ErrShiftOpenFailure
		}
	}

	// Stage 5: Open receipt
	fptr.SetParam(fptr10.LIBFPTR_PARAM_RECEIPT_TYPE, fptr10.LIBFPTR_RT_SELL)
	fptr.OpenReceipt()

	// Stage 6: Register the service or commodity
	fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, dev.cfg.ReceiptItemName)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, data.Price)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, dev.cfg.Tax)

	// Set the service tags
	// About the service provided (name and other information describing the service).
	fptr.SetParam(1212, 4)
	// Full payment, including taking into account advance payment (advance payment) at the time of transfer of the subject of calculation.
	fptr.SetParam(1214, 4)

	fptr.Registration()

	// Stage 7: Register the total
	fptr.SetParam(fptr10.LIBFPTR_PARAM_SUM, data.Price)
	fptr.ReceiptTotal()

	// Stage 8: Set the payment method
	if data.IsBankCard {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_ELECTRONICALLY)
	} else {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_CASH)
	}

	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, data.Price)
	fptr.Payment()

	// Stage 9: Close the receipt
	_ = fptr.CloseReceipt()
	if err := fptr.CheckDocumentClosed(); err != nil {
		log.Info(fptr.ErrorDescription())
		return app.ErrReceiptCloseFailure
	}
	/*
		// Stage 10: If Stage 9 failed - recover the receipt
		if !fptr.GetParamBool(fptr10.LIBFPTR_PARAM_DOCUMENT_CLOSED) {
			log.Info("Receipt can't be closed - recovering...")
			_ = fptr.CancelReceipt()
			return app.ErrReceiptCloseFailure
		}
	*/
	// Stage 11: Check the printing process
	if !fptr.GetParamBool(fptr10.LIBFPTR_PARAM_DOCUMENT_PRINTED) {
		fptr.ContinuePrint()
	}

	// Stage 12: Get fiscal data about last receipt
	fptr.SetParam(fptr10.LIBFPTR_PARAM_FN_DATA_TYPE, fptr10.LIBFPTR_FNDT_LAST_DOCUMENT)
	if err := fptr.FnQueryData(); err != nil {
		log.Info(err)
		return app.ErrUnableToGetFiscalData
	}
	log.Info("Fiscal Sign = %v", fptr.GetParamString(fptr10.LIBFPTR_PARAM_FISCAL_SIGN))
	log.Info("Fiscal Document Number = %v", fptr.GetParamInt(fptr10.LIBFPTR_PARAM_DOCUMENT_NUMBER))

	// Stage 13: Get data about unsent receipts
	fptr.SetParam(fptr10.LIBFPTR_PARAM_FN_DATA_TYPE, fptr10.LIBFPTR_FNDT_OFD_EXCHANGE_STATUS)
	if err := fptr.FnQueryData(); err != nil {
		log.Info(err)
		return app.ErrUnableToGetFiscalData
	}
	log.Printf("Unsent documents count = %v", fptr.GetParamInt(fptr10.LIBFPTR_PARAM_DOCUMENTS_COUNT))
	log.Printf("First unsent document number = %v", fptr.GetParamInt(fptr10.LIBFPTR_PARAM_DOCUMENT_NUMBER))
	log.Printf("First unsent document date = %v", fptr.GetParamDateTime(fptr10.LIBFPTR_PARAM_DATE_TIME))

	// Stage 14: Close the connection
	if err := fptr.Close(); err != nil {
		log.Info(err)
		return app.ErrCannotDisconnect
	}

	return nil
}

// NewKaznacheyFA constructs new KaznacheyFA object
func NewKaznacheyFA(mut *sync.Mutex, configSvc ConfigSvc) (*KaznacheyFA, error) {
	res := &KaznacheyFA{}
	res.mutex = mut
	for {
		var err error
		res.cfg, err = configSvc.GetConfig()

		if err == nil {
			break
		}
		log.PrintErr(err)
		time.Sleep(time.Second)
	}

	return res, nil
}
