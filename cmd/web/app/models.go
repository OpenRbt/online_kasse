package app

// CashRegisterDevice is an interface for communicating with Cash Register Device
type CashRegisterDevice interface {
	PingDevice() error
	PrintReceipt(float64, bool) error
	Start()
}

// Receipt represents generic receipt object
type Receipt struct {
	Price      float64
	IsBankCard bool
}

// NewReceipt constructs a receipt object
func NewReceipt() (*Receipt, error) {
	res := &Receipt{}
	return res, nil
}
