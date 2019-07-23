package app

import "github.com/DiaElectronics/online_kasse/cmd/web/fptr10"
import "errors"
import "strconv"

var (
	ErrCannotConnect = errors.New("Connection to KKT failed")
)

type WebApp struct {
	fptr *fptr10.IFptr
}

type CashRegisterDevice interface{
	PingDevice() error
	PrintReceipt() error
}

type Dal struct {

}

func (w *WebApp) configureKKT() {
	w.fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_KAZNACHEY_FA))
	w.fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	w.fptr.ApplySingleSettings()
}

func (w *WebApp) operatorLogin(name string, id string) {
	w.fptr.SetParam(1021, "Кассир Иванов И.")
	w.fptr.SetParam(1203, "123456789047")
	w.fptr.OperatorLogin()	
}

func (w *WebApp) openShift() error {
	w.fptr.OpenShift()
	w.fptr.CheckDocumentClosed()
	
	return nil
}

func (w *WebApp) pingKKT() bool {
	w.fptr.Open()
	isOpened := w.fptr.IsOpened()
	w.fptr.Close()
	return isOpened
}

func (w *WebApp) registerReceipt(price float64, isBankCard bool) {
	w.fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, "Мойка автомобиля");
	w.fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, price);
	w.fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1);
	w.fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, fptr10.LIBFPTR_TAX_NO);

	if isBankCard {
		w.fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_ELECTRONICALLY);
	} else {
		w.fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, fptr10.LIBFPTR_PT_CASH);
	}

	w.fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, price);
    w.fptr.Registration();
	w.fptr.Payment();
}

func (w *WebApp) PrintReceipt(price float64, isBankCard bool) error {
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

func NewWebApp () (*WebApp, error) {
	res := &WebApp{}
	return res, nil	
}



