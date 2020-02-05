package memdb

import "github.com/DiaElectronics/online_kasse/cmd/web/app"

var db *DB

var (
	r1 = app.Receipt{
		Post:       4,
		Price:      25,
		IsBankCard: false,
	}
	r2 = app.Receipt{
		Post:       3,
		Price:      2,
		IsBankCard: false,
	}
	r3 = app.Receipt{
		Post:       4,
		Price:      250,
		IsBankCard: true,
	}
)

func init() {
	db = New()
}
