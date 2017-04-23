package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"restfest/db"
	"restfest/generteststruct"
)

type copyCsv struct {
	ptr      []string
	rows     *csv.Reader
	err      error
	inter    db.InterPgx
	structer generteststruct.Weburl
}

func (t *copyCsv) Next() bool {
	if t.err != nil {
		return false
	}
	for {
		t.ptr, t.err = t.rows.Read()

		if t.err != nil {
			return false
		}

		t.err = t.inter.ConvertStoI(t.ptr)
		if t.err != nil {
			return false
		}
		if t.structer.Url != nil {
			fmt.Println(*t.structer.Url)
			break
		} else {
			fmt.Println("niller", t.structer)

		}
	}

	return true
}
func (t *copyCsv) Values() ([]interface{}, error) {
	return t.inter, t.err
}
func (t *copyCsv) Err() error {
	if t.err != io.EOF {
		return t.err
	} else {
		return nil
	}
}

func csvread() error {
	f1, err := os.Open("flat.csv")
	if err != nil {
		return err
	}

	r := csv.NewReader(f1)

	record, err := r.Read()
	if err != nil {
		return err
	}

	iterator := copyCsv{
		rows:     r,
		structer: generteststruct.Weburl{},
	}

	iterator.inter = iterator.structer.Scanner()

	fmt.Println("vor copy", record)
	copyCount, err := dbx1.CopyFrom(
		[]string{"zielweburl"},
		iterator.structer.Columns(),
		&iterator)

	fmt.Println("fertig", copyCount, err)
	if err != nil {
		return err
	}
	return nil
}
