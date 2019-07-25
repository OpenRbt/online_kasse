package dal

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// PostgresDAL represents data for connection to DB
type PostgresDAL struct {
	User     string
	Password string
	Host     string
	DataBase *pg.DB
}

// NewPostgresDAL constructs object of DB (PostgresDAL)
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
	models := (*models.Receipt)(nil)

	err := db.CreateTable(model, &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		return err
	}

	return nil
}

// CreateReceipt inserts new Receipt into DB
func (t *PostgresDAL) CreateReceipt(current *models.Receipt) (*models.Receipt, error) {
	err := t.DataBase.Insert(&current)

	if err != nil {
		return nil, err
	}

	return current, nil
}

// DeleteReceipt deletes current Receipt by ID
func (t *PostgresDAL) DeleteReceipt(current *models.Receipt) (*models.Receipt, error) {
	// TO DO: add ID field to the model of Receipt
	err := t.DataBase.Delete(&current)

	if err != nil {
		return nil, err
	}

	return current, nil
}

// GetProcessed returns a list of transfered receipts
func (t *PostgresDAL) GetProcessed(current *models.GetData) (*models.ReceiptList, error) {
	var receipts []models.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return models.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetUnprocessed returns a list of transfered receipts
func (t *PostgresDAL) GetUnprocessed(current *models.GetData) (*models.ReceiptList, error) {
	var receipts []models.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return models.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetCash returns a list of transfered receipts
func (t *PostgresDAL) GetCash(current *models.GetData) (*models.ReceiptList, error) {
	var receipts []models.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return models.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetBankCards returns a list of transfered receipts
func (t *PostgresDAL) GetBankCards(current *models.GetData) (*models.ReceiptList, error) {
	var receipts []models.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return models.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}

// GetPrices returns a list of transfered receipts
func (t *PostgresDAL) GetPrices(current *models.GetData) (*models.ReceiptList, error) {
	var receipts []models.Receipt

	err := t.DataBase.Model(&receipts).Where("receipt_id = ?", current.ReceiptID).Select()
	if err != nil {
		return nil, err
	}

	return models.ReceiptList{Receipts: receipts, Total: len(receipts)}, nil
}
