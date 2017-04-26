package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"restfest/generteststruct"
)

type copyCsv struct {
	ptr  []string
	rows *csv.Reader
	generteststruct.IterLos
	anz int
}

func (t *copyCsv) Next() bool {
	if t.Errc != nil {
		return false
	}
	for {
		t.ptr, t.Errc = t.rows.Read()

		if t.Errc != nil {
			return false
		}

		t.Errc = t.Inter.ConvertStoI(t.ptr)
		if t.Errc != nil {
			return false
		}
		if t.Los.L_iban != nil {
			//			fmt.Println(*t.structer.L_iban)
			t.anz++
			if t.anz%1000 == 0 {
				fmt.Println("read", t.anz)
			}
			return true
		}
	}

}

func csvread() error {
	f1, err := os.Open("flat.csv")
	if err != nil {
		return err
	}

	iterator := copyCsv{
		rows: csv.NewReader(f1),
	}

	iterator.StartCopy("ziellos", dbx1, &iterator, nil)
	return nil
}
