package dal

import (
	"fmt"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Receipt represents generic Receipt object in DAL
type Receipt struct {
	Id          int64
	Price       float64
	IsBankCard  bool
	IsProcessed bool
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
	fmt.Println(q.FormattedQuery())
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
	target.IsBankCard = current.IsBankCard

	err := t.DataBase.Insert(&target)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return current, nil
}

// DeleteByID deletes specified Receipt by ID
func (t *PostgresDAL) DeleteByID(ID int64) (int64, error) {
	var target Receipt
	target.Id = ID
	target.IsProcessed = false

	err := t.DataBase.Delete(&target)

	if err != nil {
		return -1, err
	}

	return ID, nil
}

// UpdateStatus changes IsProcessed field to true
func (t *PostgresDAL) UpdateStatus(current *app.Receipt) (bool, error) {
	var target Receipt
	target.Id = current.Id
	target.IsProcessed = true
	target.Price = current.Price
	target.IsBankCard = current.IsBankCard

	err := t.DataBase.Update(&target)

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetProcessedOnly returns a list of processed (transfered) Receipts
func (t *PostgresDAL) GetProcessedOnly(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_processed = ?", true).Where("id >= ?", current.LastId).Limit(current.Limit).Select()

	if err != nil {
		return nil, err
	}

	for _, element := range foundReceipts {
		var newReceipt app.Receipt

		newReceipt.Id = element.Id
		newReceipt.Price = element.Price
		newReceipt.IsBankCard = element.IsBankCard

		convertedReceipts = append(convertedReceipts, newReceipt)
	}

	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetUnprocessedOnly returns a list of unprocessed (untransfered) Receipts
func (t *PostgresDAL) GetUnprocessedOnly(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_processed = 0").Where("id >= ?", current.LastId).Limit(current.Limit).Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, element := range foundReceipts {
		var newReceipt app.Receipt

		newReceipt.Id = element.Id
		newReceipt.Price = element.Price
		newReceipt.IsBankCard = element.IsBankCard

		convertedReceipts = append(convertedReceipts, newReceipt)
	}

	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetByPrice returns a list of Receipts by specified price (for instance: 50 or 100 RUB)
func (t *PostgresDAL) GetByPrice(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("price = ?", current.Price).Where("id >= ?", current.LastId).Limit(current.Limit).Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, element := range foundReceipts {
		var newReceipt app.Receipt

		newReceipt.Id = element.Id
		newReceipt.Price = element.Price
		newReceipt.IsBankCard = element.IsBankCard

		convertedReceipts = append(convertedReceipts, newReceipt)
	}

	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetWithBankCards returns a list of Receipts paid by Bank Cards only
func (t *PostgresDAL) GetWithBankCards(current app.QueryData) (*app.ReceiptList, error) {
	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_bank_card = 1").Where("id >= ?", current.LastId).Limit(current.Limit).Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, element := range foundReceipts {
		var newReceipt app.Receipt

		newReceipt.Id = element.Id
		newReceipt.Price = element.Price
		newReceipt.IsBankCard = element.IsBankCard

		convertedReceipts = append(convertedReceipts, newReceipt)
	}

	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}

// GetWithCash returns a list of Receipts paid by Cash only
func (t *PostgresDAL) GetWithCash(current app.QueryData) (*app.ReceiptList, error) {

	var foundReceipts []Receipt
	var convertedReceipts []app.Receipt

	err := t.DataBase.Model(&foundReceipts).Where("is_bank_card = 0").Where("id >= ?", current.LastId).Limit(current.Limit).Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, element := range foundReceipts {
		var newReceipt app.Receipt

		newReceipt.Id = element.Id
		newReceipt.Price = element.Price
		newReceipt.IsBankCard = element.IsBankCard

		convertedReceipts = append(convertedReceipts, newReceipt)
	}

	return &app.ReceiptList{Receipts: convertedReceipts, Total: len(convertedReceipts)}, nil
}
