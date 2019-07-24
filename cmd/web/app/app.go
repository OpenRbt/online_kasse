package app

import "github.com/DiaElectronics/online_kasse/cmd/web/fptr10"
import "errors"
import "strconv"

var (
	ErrCannotConnect = errors.New("Connection to Cash Register Device failed")
)

type Application struct {

}

func (a *Application) PingDevice() error {
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

	return nil
}

func (a *Application) PrintReceipt(price float64, isBankCard bool) error {
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
	
	//fptr.SetParam(fptr10.LIBFPTR_PARAM_REPORT_TYPE, fptr10.LIBFPTR_RT_CLOSE_SHIFT)
	//fptr.Report()
	
	fptr.OpenShift()
	
	fptr.SetParam(fptr10.LIBFPTR_PARAM_RECEIPT_TYPE, fptr10.LIBFPTR_RT_SELL)
	fptr.OpenReceipt()
	
	fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, "Мойка автомобиля")
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, price)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1)
	fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, fptr10.LIBFPTR_TAX_NO);
	fptr.Registration()
	
	if isBankCard {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_ELECTRONICALLY);
	} else {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_CASH);
	}
	
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, price);
	fptr.Payment();
	
	fptr.CloseReceipt()
	fptr.CheckDocumentClosed()

	fptr.Close()
	
	fptr.Destroy()

	return nil
} 

func NewApplication () (*Application, error) {
	res := &Application{}
	return res, nil	
}

func (a *Application) Start() {
	// TO DO: start goroutine with data processing from DB
}



