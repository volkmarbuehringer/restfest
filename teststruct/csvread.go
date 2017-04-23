package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"restfest/db"
	"restfest/generteststruct"

	"github.com/jackc/pgx"
)

type baseCopy struct {
	err      error
	inter    db.InterPgx
	structer generteststruct.Weburl
}

type copyCsv struct {
	ptr  []string
	rows *csv.Reader
	baseCopy
}

func (t *baseCopy) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource) error {

	t.inter = t.structer.Scanner()

	fmt.Println("vor copy", tab)
	copyCount, err := con.CopyFrom(
		[]string{tab},
		t.structer.Columns(),
		tt)

	fmt.Println("fertig", copyCount, err)
	if err != nil {
		return err
	}
	return nil
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
func (t *baseCopy) Values() ([]interface{}, error) {
	return t.inter, t.err
}
func (t baseCopy) Err() error {
	if t.err != io.EOF {
		return t.err
	}
	return nil
}

func csvread() error {
	f1, err := os.Open("flat.csv")
	if err != nil {
		return err
	}

	r := csv.NewReader(f1)

	_, err = r.Read()
	if err != nil {
		return err
	}

	iterator := copyCsv{
		rows:     r,
		baseCopy: baseCopy{structer: generteststruct.Weburl{}},
	}

	iterator.StartCopy("zielweburl", dbx1, &iterator)
	return nil
}
