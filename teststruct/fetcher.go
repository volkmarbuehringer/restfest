package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"restfest/db"
	"restfest/generteststruct"

	log15 "gopkg.in/inconshreveable/log15.v2"
)

type scanner struct {
	generteststruct.BaseCopyLos
}

func fetcher() error {

	f1, err := os.Create("flat.csv")
	if err != nil {
		log15.Crit("DBFehler", "create", err)
		return err
	}
	w := csv.NewWriter(f1)
	w.Comma = ';'
	rows, err := dbx3.Query(generteststruct.SQLLos(db.GenSelectAll1), 30000000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return err
	}
	defer rows.Close()

	var iter scanner

	iter.NewCopy(rows)

	for iter.Next() {
		if iter.Los.L_iban != nil {
			fmt.Println(*iter.Los.L_iban)
			record, err := iter.ValuesString()
			if err != nil {
				return err
			}
			if err := w.Write(record); err != nil {
				return err
			}

		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}
	return rows.Err()
}
