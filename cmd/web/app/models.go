package app

// Receipt represents generic Receipt object
type Receipt struct {
	Price      float64
	IsBankCard bool
}

// ReceiptList represents list of Receipts
type ReceiptList struct {
	Receipts []Receipt
	Total    int
}

// QueryData represents object for transfering query config data to DAL
type QueryData struct {
	Limit     int
	LastID    int
	ReceiptID int
}

// NewReceipt constructs a Receipt object
func NewReceipt() (*Receipt, error) {
	res := &Receipt{}
	return res, nil
}
