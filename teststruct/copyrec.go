package main

import (
	"fmt"
	"restfest/db"
	"restfest/generteststruct"

	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

type weburlScan struct {
	ptr   generteststruct.Weburl
	rows  *pgx.Rows
	inter db.InterPgx
}

func (t *weburlScan) Next() bool {
	var ok bool
	for {
		ok = t.rows.Next()
		if !ok {
			break
		}
		t.rows.Scan(t.inter...)
		if t.ptr.Url != nil {
			fmt.Println(*t.ptr.Url)
			break
		}
	}
	return ok
}
func (t *weburlScan) Values() ([]interface{}, error) {
	return t.inter, t.rows.Err()
}
func (t *weburlScan) Err() error {
	return t.rows.Err()
}
func copyer() error {
	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return err
	}
	defer rows.Close()

	iterator := weburlScan{
		ptr:  generteststruct.Weburl{},
		rows: rows,
	}
	iterator.inter = iterator.ptr.Scanner()
	fmt.Println("vor copy")
	copyCount, err := dbx1.CopyFrom(
		[]string{"zielweburl"},
		iterator.ptr.Columns(),
		&iterator)

	fmt.Println("fertig", copyCount, err)
	return err
}
