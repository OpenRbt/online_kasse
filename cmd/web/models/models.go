package models

// CashRegisterDevice is an interface for communicating with Cash Register Device
type CashRegisterDevice interface {
	PingDevice() error
	PrintReceipt() error
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
