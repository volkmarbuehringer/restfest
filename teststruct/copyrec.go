package main

import (
	"fmt"
	"restfest/db"
	"restfest/generteststruct"

	log15 "gopkg.in/inconshreveable/log15.v2"
)

type weburlScan struct {
	generteststruct.BaseCopyLos
}

func (t *weburlScan) Next() bool {
	for t.BaseCopyLos.Next() {
		if t.Structer.L_iban != nil {
			fmt.Println(*t.Structer.L_iban)
			return true
		}
	}
	return false
}

func copyer() error {
	rows, err := dbx.Query(generteststruct.SQLLos(db.GenSelectAll1), 40000000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return err
	}
	defer rows.Close()

	iterator := weburlScan{
		BaseCopyLos: generteststruct.BaseCopyLos{Structer: generteststruct.Los{}, Rows: rows},
	}

	return iterator.StartCopy("ziellos", dbx2, &iterator)

}
