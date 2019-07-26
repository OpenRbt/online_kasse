package dal

import (
	"fmt"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

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
	if db != nil {
		fmt.Println("PostgreSQL connected")
	}

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
	for _, model := range []interface{}{(*app.Receipt)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// Create inserts new Receipt into DB
func (t *PostgresDAL) Create(current *app.Receipt) (*app.Receipt, error) {
	err := t.DataBase.Insert(&current)

	if err != nil {
		return nil, err
	}

	return current, nil
}

// DeleteByID deletes specified Receipt by ID
func (t *PostgresDAL) DeleteByID(current *app.Receipt) (*app.Receipt, error) {
	// TO DO: add ID field to the model of Receipt
	err := t.DataBase.Delete(&current)

	if err != nil {
		return nil, err
	}

	return current, nil
}

// GetProcessedOnly returns a list of processed (transfered) Receipts
func (t *PostgresDAL) GetProcessedOnly(current *app.QueryData) (*app.ReceiptList, error) {
	var receipts []app.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return &app.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetUnprocessedOnly returns a list of unprocessed (untransfered) Receipts
func (t *PostgresDAL) GetUnprocessedOnly(current *app.QueryData) (*app.ReceiptList, error) {
	var receipts []app.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return &app.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetByPrice returns a list of Receipts by specified price (for instance: 50 or 100 RUB)
func (t *PostgresDAL) GetByPrice(current *app.QueryData) (*app.ReceiptList, error) {
	var receipts []app.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return &app.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetWithBankCards returns a list of Receipts paid by Bank Cards only
func (t *PostgresDAL) GetWithBankCards(current *app.QueryData) (*app.ReceiptList, error) {
	var receipts []app.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return &app.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetWithCash returns a list of Receipts paid by Cash only
func (t *PostgresDAL) GetWithCash(current *app.QueryData) (*app.ReceiptList, error) {
	var receipts []app.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return &app.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}
