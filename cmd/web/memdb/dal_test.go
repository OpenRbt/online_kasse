package memdb

import (
	"testing"

	"github.com/DiaElectronics/online_kasse/cmd/web/app"
	"github.com/powerman/check"
)

func TestSmoke(tt *testing.T) {
	t := check.T(tt)
	_, err := db.Create(&r1)
	t.Nil(err)
	r1.ID = 1
	_, err = db.Create(&r2)
	t.Nil(err)
	r2.ID = 2
	_, err = db.Create(&r3)
	t.Nil(err)
	r3.ID = 3

	res, err := db.GetUnprocessedOnly(app.QueryData{Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r1, r2, r3},
		Total:    3,
	})

	_, err = db.UpdateStatus(r1)
	t.Nil(err)

	res, err = db.GetUnprocessedOnly(app.QueryData{Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r2, r3},
		Total:    2,
	})

	res, err = db.GetProcessedOnly(app.QueryData{Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r1},
		Total:    1,
	})

	res, err = db.GetWithBankCards(app.QueryData{Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r3},
		Total:    1,
	})

	res, err = db.GetWithCash(app.QueryData{Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r1, r2},
		Total:    2,
	})

	res, err = db.GetByPost(app.QueryData{Post: 4, Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r1, r3},
		Total:    2,
	})

	_, err = db.DeleteByID(1)
	t.Nil(err)
	_, err = db.DeleteByID(3)
	t.Nil(err)

	res, err = db.GetWithCash(app.QueryData{Limit: 5, LastID: 0})
	t.Nil(err)
	t.DeepEqual(res, &app.ReceiptList{
		Receipts: []app.Receipt{r2},
		Total:    1,
	})

	t.Equal(db.Info(), "memdb")
}
