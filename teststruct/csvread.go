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
	generteststruct.BaseCopyLos
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
			break
		} else {
			//			fmt.Println("niller")

		}
	}

	return true
}

func csvread() error {
	f1, err := os.Open("flat.csv")
	if err != nil {
		return err
	}

	r := csv.NewReader(f1)

	iterator := copyCsv{
		rows:        r,
		BaseCopyLos: generteststruct.BaseCopyLos{Los: generteststruct.Los{}},
	}

	iterator.StartCopy("ziellos", dbx1, &iterator)
	return nil
}
