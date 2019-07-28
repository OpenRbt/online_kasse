package app

// Receipt represents generic Receipt object
type Receipt struct {
	ID         int64
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
	Limit  int
	LastID int
	Price  int
}

// NewReceipt constructs a Receipt object
func NewReceipt() *Receipt {
	return &Receipt{}
}
