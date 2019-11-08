package dal

import (
	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Receipt represents generic Receipt object in DAL
type Receipt struct {
	ID          int64
	Post        int64
	Price       float64
	IsBankCard  int8
	IsProcessed int8
}

// PostgresDAL represents data for connection to Data base
type PostgresDAL struct {
	User     string
	Password string
	Host     string
	DataBase *pg.DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
}

// NewPostgresDAL constructs object of PostgresDAL
func NewPostgresDAL(user string, password string, host string) (*PostgresDAL, error) {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Addr:     host,
	})
	db.AddQueryHook(dbLogger{})

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	res := &PostgresDAL{
		User:     user,
		Password: password,
		Host:     host,
		DataBase: db}

	return res, nil
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

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Receipt)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Create inserts new Receipt into DB
func (t *PostgresDAL) Create(current *app.Receipt) (*app.Receipt, error) {
	var target Receipt
	target.Price = current.Price
	target.Post = current.Post

	if current.IsBankCard {
		target.IsBankCard = 1
	} else {
		target.IsBankCard = -1
	}

	target.IsProcessed = -1

	err := t.DataBase.Insert(&target)

	if err != nil {
		return nil, err
	}

	return current, nil
}

// DeleteByID deletes specified Receipt by ID
func (t *PostgresDAL) DeleteByID(ID int64) (int64, error) {
	var target Receipt
	target.ID = ID

	err := t.DataBase.Delete(&target)

	if err != nil {
		return -1, err
	}

	return ID, nil
}

// UpdateStatus changes IsProcessed field to true
func (t *PostgresDAL) UpdateStatus(current app.Receipt) (bool, error) {
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

	err := t.DataBase.Update(&target)

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetProcessedOnly returns a list of processed (transfered) Receipts
func (t *PostgresDAL) GetProcessedOnly(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_processed = 1").Where("id > ?", current.LastID).Limit(current.Limit).Select()

	if err != nil {
		return nil, err
	}

	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetUnprocessedOnly returns a list of unprocessed (untransfered) Receipts
func (t *PostgresDAL) GetUnprocessedOnly(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_processed = -1").Where("id >= ?", current.LastID).Limit(current.Limit).Select()

	if err != nil {
		return nil, err
	}

	convertedReceipts := makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetWithBankCards returns a list of Receipts paid by Bank Cards only
func (t *PostgresDAL) GetWithBankCards(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_bank_card = 1").Where("id >= ?", current.LastID).Limit(current.Limit).Select()

	if err != nil {
		return nil, err
	}

	convertedReceipts = makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetWithCash returns a list of Receipts paid by Cash only
func (t *PostgresDAL) GetWithCash(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_bank_card = -1").Where("id >= ?", current.LastID).Limit(current.Limit).Select()

	if err != nil {
		return nil, err
	}

	convertedReceipts = makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetByPost returns a list of Receipts by specified post number
func (t *PostgresDAL) GetByPost(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("post = ?", current.Post).Where("id >= ?", current.LastID).Limit(current.Limit).Select()

	if err != nil {
		return nil, err
	}

	convertedReceipts = makeAppReceiptSlice(foundReceipts)
	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}
