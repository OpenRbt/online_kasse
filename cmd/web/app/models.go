package app

type CashRegisterDevice interface {
	PingDevice() error
	PrintReceipt() error
}

type Receipt struct {
	Price float64
	IsBankCard bool
}

func NewReceipt () (*Receipt, error) {
	res := &Receipt{}
	return res, nil	
}