package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"restfest/db"
	"restfest/teststruct/gener"

	"github.com/Sirupsen/logrus"
)

func fetcher() error {

	f1, err := os.Create("flat.csv")
	if err != nil {
		logrus.Errorf("DBFehler %v", err)
		return err
	}
	w := csv.NewWriter(f1)
	w.Comma = ';'
	params := new(gener.LosParams)
	rows, err := dbx3.Query(params.SQL(db.GenSelectAll1), 30000000, 0)
	if err != nil {
		logrus.Errorf("DBFehler %v", err)
		return err
	}
	defer rows.Close()

	var iter gener.IterLos

	iter.NewCopy(rows)

	for iter.Next() {
		if iter.Los.L_iban != nil { //check values in struct
			fmt.Println(*iter.Los.L_iban)

		} else { //replace values in struct
			iter.Los.L_iban = db.String("willi")
		}
		record, err := iter.ValuesString()
		if err != nil {
			return err
		}
		if err := w.Write(record); err != nil {
			return err
		}

	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}
	return rows.Err()
}
