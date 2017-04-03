package main

import (
	"fmt"
	"restfest/db"

	"restfest/gener"
)

func main() {

	sql := gener.SQLGutschein("gutschein", db.GenSelect)
	rows, err := db.DB.Query("select " + sql[0] + " from gutschein")
	if err != nil {
		return
	}

	defer rows.Close()
	for anz := 0; rows.Next(); anz++ {
		arr, ts := gener.ScannerTGutschein()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		if ts.Auf_partner.Valid {
			fmt.Printf("%d %s \n", ts.Auf_id.Int64, ts.Auf_partner.String)
		} else {
			fmt.Printf("%d null \n", ts.Auf_id.Int64)
		}
		//fmt.Printf("%v %T \n", ts, ts)
	}
}
