package app

// CashRegisterDevice is an interface for communicating with Cash Register Device
type CashRegisterDevice interface {
	PingDevice() error
	RegisterReceipt()
	ResetShift()
	PrintReceipt(float64, bool) error
	Start()
}

// Receipt represents generic receipt object
type Receipt struct {
	Price      float64
	IsBankCard bool
}

// ReceiptList represents list of generic receipt objects
type ReceiptList struct {
}

// GetData represents object for transfering config data to DAL
type GetData struct {
}

// NewReceipt constructs a receipt object
func NewReceipt() (*Receipt, error) {
	res := &Receipt{}
	return res, nil
}
