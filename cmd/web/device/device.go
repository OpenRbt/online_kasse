package device

import (
	"strconv"
	"sync"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/DiaElectronics/online_kasse/cmd/web/fptr10"
	"github.com/powerman/structlog"
)

var log = structlog.New()

// KaznacheyFA representes object of Device, connected by USB
type KaznacheyFA struct {
	mutex sync.Mutex
}

// ResetShift sends signal to Device, which will close current shift and will open new one
//nolint
func (dev *KaznacheyFA) ResetShift() error {
	dev.mutex.Lock()
	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()
	if !fptr.IsOpened() {
		return app.ErrCannotConnect
	}

	fptr.SetParam(1021, "Кассир Канатников Александр")
	fptr.SetParam(1203, "123456789047")
	fptr.OperatorLogin()

	fptr.SetParam(fptr10.LIBFPTR_PARAM_REPORT_TYPE, fptr10.LIBFPTR_RT_CLOSE_SHIFT)
	fptr.Report()

	fptr.OpenShift()

	fptr.Close()

	fptr.Destroy()
	dev.mutex.Unlock()

	return nil
}

// PingDevice checks connection to the Device
//nolint
func (dev *KaznacheyFA) PingDevice() error {
	dev.mutex.Lock()
	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()
	if !fptr.IsOpened() {
		fptr.Destroy()
		return app.ErrCannotConnect
	}
	fptr.Close()

	fptr.Destroy()

	dev.mutex.Unlock()
	return nil
}

// PrintReceipt sends Receipt to the Device driver
//nolint
func (dev *KaznacheyFA) PrintReceipt(data app.Receipt) error {
	dev.mutex.Lock()
	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()
	if !fptr.IsOpened() {
		return app.ErrCannotConnect
	}

	log.Info("Connection to kasse opened")

	fptr.SetParam(1021, "Кассир Канатников Александр")
	fptr.SetParam(1203, "123456789047")
	fptr.OperatorLogin()

	fptr.OpenShift()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.SetParam(fptr10.LIBFPTR_PARAM_RECEIPT_TYPE, fptr10.LIBFPTR_RT_SELL)
	fptr.OpenReceipt()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, "Мойка автомобиля")
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, data.Price)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, fptr10.LIBFPTR_TAX_NO)
	fptr.Registration()

	if data.IsBankCard {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_ELECTRONICALLY)
	} else {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_CASH)
	}

	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, data.Price)
	fptr.Payment()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.CloseReceipt()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.CheckDocumentClosed()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.Close()

	fptr.Destroy()
	dev.mutex.Unlock()

	return nil
}

// NewKaznacheyFA constructs new KaznacheyFA object
func NewKaznacheyFA(mut sync.Mutex) (*KaznacheyFA, error) {
	res := &KaznacheyFA{}
	res.mutex = mut

	fptr := fptr10.New()

	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_ATOL_AUTO))
	fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	fptr.ApplySingleSettings()

	fptr.Open()

	log.Info("Connection to kasse opened")

	fptr.SetParam(fptr10.LIBFPTR_PARAM_REPORT_TYPE, fptr10.LIBFPTR_RT_CLOSE_SHIFT)
	fptr.Report()

	fptr.CheckDocumentClosed()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.SetParam(1021, "Пост 1")
	fptr.SetParam(1203, "123456789047")
	fptr.OperatorLogin()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.OpenShift()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.SetParam(fptr10.LIBFPTR_PARAM_RECEIPT_TYPE, fptr10.LIBFPTR_RT_SELL)
	fptr.OpenReceipt()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, "Мойка автомобиля")
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, 10)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, fptr10.LIBFPTR_TAX_NO)
	fptr.Registration()

	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_CASH)

	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, 10)
	fptr.Payment()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.CloseReceipt()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.CheckDocumentClosed()

	log.Info(fptr.ErrorCode())
	log.Info(fptr.ErrorDescription())

	fptr.Close()

	fptr.Destroy()

	return res, nil
}
