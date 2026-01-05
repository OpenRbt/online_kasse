package memdb

import "github.com/OpenRbt/online_kasse/cmd/web/app"

var db *DB

var (
	r1 = app.Receipt{
		Post: 4,
		Cash: 25,
	}
	r2 = app.Receipt{
		Post: 3,
		Cash: 2,
	}
	r3 = app.Receipt{
		Post:           4,
		Electronically: 250,
	}
)

func init() {
	db = New()
}
