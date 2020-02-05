package memdb

import (
	"sync"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
)

// Receipt represents generic Receipt object in DAL
type Receipt struct {
	ID          int64
	Post        int64
	Price       float64
	IsBankCard  int8
	IsProcessed int8
}

type DB struct {
	receipt   []Receipt
	mutex     sync.Mutex
	receiptID int64
}

// New DB
func New() *DB {
	return &DB{receipt: []Receipt{}}
}

func makeAppReceipt(from Receipt) app.Receipt {
	var appReceipt app.Receipt

	appReceipt.Price = from.Price
	appReceipt.ID = from.ID
	appReceipt.Post = from.Post

	if from.IsBankCard == 1 {
		appReceipt.IsBankCard = true
	} else if from.IsBankCard == -1 {
		appReceipt.IsBankCard = false
	}

	return appReceipt
}

func makeAppReceiptSlice(from []Receipt) []app.Receipt {
	var appReceipts []app.Receipt

	for _, element := range from {
		newReceipt := makeAppReceipt(element)
		appReceipts = append(appReceipts, newReceipt)
	}

	return appReceipts
}

// Create inserts new Receipt into DB
func (t *DB) Create(current *app.Receipt) (*app.Receipt, error) {
	var target Receipt
	target.Price = current.Price
	target.Post = current.Post

	if current.IsBankCard {
		target.IsBankCard = 1
	} else {
		target.IsBankCard = -1
	}

	target.IsProcessed = -1
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.receiptID++
	target.ID = t.receiptID
	t.receipt = append(t.receipt, target)
	return current, nil
}

// DeleteByID deletes specified Receipt by ID
func (t *DB) DeleteByID(ID int64) (int64, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	i := t.findByID(ID)
	if i < 0 {
		return -1, app.ErrNotFound
	}
	t.receipt = append(t.receipt[:i], t.receipt[i+1:]...)
	return ID, nil
}

func (t *DB) findByID(id int64) int64 {
	for i, _ := range t.receipt {
		if t.receipt[i].ID == id {
			return int64(i)
		}
	}
	return -1
}

// UpdateStatus changes IsProcessed field to true
func (t *DB) UpdateStatus(current app.Receipt) (bool, error) {
	var target Receipt
	target.ID = current.ID
	target.IsProcessed = 1
	target.Price = current.Price
	target.Post = current.Post

	if current.IsBankCard {
		target.IsBankCard = 1
	} else {
		target.IsBankCard = -1
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	i := t.findByID(target.ID)
	if i < 0 {
		return false, app.ErrNotFound
	}
	t.receipt[i] = target
	return true, nil
}

// GetProcessedOnly returns a list of processed (transfered) Receipts
func (t *DB) GetProcessedOnly(current app.QueryData) (*app.ReceiptList, error) {
	if current.Limit < 1 {
		return &app.ReceiptList{}, nil
	}

	var foundReceipts []Receipt
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for i, _ := range t.receipt {
		if t.receipt[i].ID > int64(current.LastID) && t.receipt[i].IsProcessed == 1 {
			foundReceipts = append(foundReceipts, t.receipt[i])
			if len(foundReceipts) == current.Limit {
				break
			}
		}
	}
	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetUnprocessedOnly returns a list of unprocessed (untransfered) Receipts
func (t *DB) GetUnprocessedOnly(current app.QueryData) (*app.ReceiptList, error) {
	if current.Limit < 1 {
		return &app.ReceiptList{}, nil
	}

	var foundReceipts []Receipt
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for i, _ := range t.receipt {
		if t.receipt[i].ID > int64(current.LastID) && t.receipt[i].IsProcessed == -1 {
			foundReceipts = append(foundReceipts, t.receipt[i])
			if len(foundReceipts) == current.Limit {
				break
			}
		}
	}

	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetWithBankCards returns a list of Receipts paid by Bank Cards only
func (t *DB) GetWithBankCards(current app.QueryData) (*app.ReceiptList, error) {
	if current.Limit < 1 {
		return &app.ReceiptList{}, nil
	}

	var foundReceipts []Receipt
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for i, _ := range t.receipt {
		if t.receipt[i].ID > int64(current.LastID) && t.receipt[i].IsBankCard == 1 {
			foundReceipts = append(foundReceipts, t.receipt[i])
			if len(foundReceipts) == current.Limit {
				break
			}
		}
	}

	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetWithCash returns a list of Receipts paid by Cash only
func (t *DB) GetWithCash(current app.QueryData) (*app.ReceiptList, error) {
	if current.Limit < 1 {
		return &app.ReceiptList{}, nil
	}

	var foundReceipts []Receipt
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for i, _ := range t.receipt {
		if t.receipt[i].ID > int64(current.LastID) && t.receipt[i].IsBankCard == -1 {
			foundReceipts = append(foundReceipts, t.receipt[i])
			if len(foundReceipts) == current.Limit {
				break
			}
		}
	}

	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetByPost returns a list of Receipts by specified post number
func (t *DB) GetByPost(current app.QueryData) (*app.ReceiptList, error) {
	if current.Limit < 1 {
		return &app.ReceiptList{}, nil
	}

	var foundReceipts []Receipt
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for i, _ := range t.receipt {
		if t.receipt[i].ID > int64(current.LastID) && t.receipt[i].Post == int64(current.Post) {
			foundReceipts = append(foundReceipts, t.receipt[i])
			if len(foundReceipts) == current.Limit {
				break
			}
		}
	}

	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// Info returns database information
func (t *DB) Info() string {
	return "memdb"
}
