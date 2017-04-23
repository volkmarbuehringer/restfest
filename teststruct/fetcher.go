package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"restfest/db"
	"restfest/generteststruct"

	log15 "gopkg.in/inconshreveable/log15.v2"
)

func fetcher() error {

	f1, err := os.Create("flat.csv")
	if err != nil {
		log15.Crit("DBFehler", "create", err)
		return err
	}
	w := csv.NewWriter(f1)

	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return err
	}
	defer rows.Close()
	stru := new(generteststruct.Weburl)
	arrp := stru.Scanner()
	for anz := 0; rows.Next(); anz++ {

		if err = rows.Scan(arrp...); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return err
		}
		if stru.Url != nil {
			fmt.Println(*stru.Url)
		}

		record, err := arrp.ConvertItoS()
		if err != nil {
			return err
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}
	return rows.Err()
}
