package main

import (
	"fmt"
	"restfest/db"
	"restfest/generteststruct"

	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

type weburlScan struct {
	rows *pgx.Rows
	baseCopy
}

func (t *weburlScan) Next() bool {

	var ok bool
	for {
		ok = t.rows.Next()
		if !ok {
			break
		}
		t.rows.Scan(t.inter...)
		if t.structer.Url != nil {
			fmt.Println(*t.structer.Url)
			break
		}
	}
	t.err = t.rows.Err()
	if t.err != nil {
		return false
	}
	return ok
}

func copyer() error {
	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return err
	}
	defer rows.Close()

	iterator := weburlScan{
		baseCopy: baseCopy{structer: generteststruct.Weburl{}},
		rows:     rows,
	}

	return iterator.StartCopy("zielweburl", dbx2, &iterator)

}
