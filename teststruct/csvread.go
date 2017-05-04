package main

import (
	"encoding/csv"
	"fmt"
	"os"
	gener "restfest/restspec/gener"
)

type copyCsv struct {
	ptr  []string
	rows *csv.Reader
	gener.IterLos
	anz int
}

func (t *copyCsv) Next() bool {

	for {
		t.ptr, t.Errc = t.rows.Read()

		if t.Errc != nil {
			return false
		}
		//	fmt.Println("string", len(t.ptr), t.ptr)
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
	w := csv.NewReader(f1)
	w.Comma = ';'
	iterator := copyCsv{
		rows: w,
	}

	count, err := iterator.StartCopy("ziellos", dbx1, &iterator, nil)
	fmt.Println("nach copy", count)
	if err != nil {
		return err
	}

	return nil
}
