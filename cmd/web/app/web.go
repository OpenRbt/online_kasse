package app

import "atol.ru/drivers10/fptr"
import "errors"

var (
	ErrCannotConnect = errors.new("Connection to KKT failed")
)

type WebApp struct {
	fptr IFptr
}

type Dal struct {

}

func (w *WebApp) configureKKT() {
	w.fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_MODEL, strconv.Itoa(fptr10.LIBFPTR_MODEL_KAZNACHEY_FA))
	w.fptr.SetSingleSetting(fptr10.LIBFPTR_SETTING_PORT, strconv.Itoa(fptr10.LIBFPTR_PORT_USB))
	w.fptr.ApplySingleSettings()
}

func (w *WebApp) operatorLogin(string name, string id) {
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
	fptr.SetParam(fptr10.LIBFPTR_PARAM_COMMODITY_NAME, "Мойка автомобиля");
	fptr.SetParam(fptr10.LIBFPTR_PARAM_PRICE, price);
	fptr.SetParam(fptr10.LIBFPTR_PARAM_QUANTITY, 1);
	fptr.SetParam(fptr10.LIBFPTR_PARAM_TAX_TYPE, LIBFPTR_TAX_NO);

    if isBankCard == false {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, LIBFPTR_PT_CASH);
	}
    else {
		fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_TYPE, LIBFPTR_PT_ELECTRONICALLY);
	}

	fptr.SetParam(fptr10.LIBFPTR_PARAM_PAYMENT_SUM, price);
    fptr.Registration();
	fptr.Payment();
}

func (w *WebApp) PrintReceipt(price float64, isBankCard bool) error {
	fptr := fptr10.New()

	w.configureKKT()

	isAvailable := pingKKT()
	if isAvailable == false {
		return ErrCannotConnect
	}

	fptr.Open()
	w.operatorLogin("Ivan Pupkin", "123456789")

	w.registerReceipt(price, isBankCard)

	fptr.Close()
	fptr.Destroy()

	return nil
} 

func NewWebApp () (*WebApp, error) {
	res := &WebApp{}
	return res, nil	
}



