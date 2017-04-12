package main

import (
	"fmt"
	"os"
	"restfest/db"
	"restfest/gener"
	"strconv"

	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

func main() {

	limiter, _ := strconv.Atoi(os.Args[1])
	rows, err := db.DBx.Query(gener.SQLGutschein(db.GenSelectAll), limiter, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	store := make(map[int32]*gener.Gutschein, 0)
	defer rows.Close()
	for anz := 0; rows.Next() && anz < limiter; anz++ {
		if err = rows.Err(); err != nil {
			log15.Crit("DBFehler", "get", err)
			return
		}

		arr, ts := gener.ScannerTGutschein()
		if err = rows.Scan(arr...); err != nil {
			log15.Crit("DBFehler", "get", err)
			return
		}
		store[*ts.Aug_id] = ts

		if ts.Auf_anzahl != nil {
			tt, _ := strconv.Atoi(*ts.Auf_anzahl)
			if tt > 1 && ts.Auf_upd_uid != nil && ts.Auf_bemerkung != nil {
				fmt.Printf("grösser %s %s\n", *ts.Auf_upd_uid, *ts.Auf_bemerkung)
			}

		}

		if ts.Auf_partner != nil && len(*ts.Auf_partner) > 0 {
			fmt.Printf("%d %s \n", ts.Auf_id, *ts.Auf_partner)
		} else {
			fmt.Printf("%d null \n", ts.Auf_id)
		}
	}
	defer db.DBx.Close()

	rower := make([][]interface{}, 0)
	for _, t := range store {
		rower = append(rower, []interface{}{t.Aug_id, t.Auf_bemerkung, t.Auf_upd_uid})
	}
	fmt.Println("vor copy", len(rower))
	if copyCount, err := db.DBx.CopyFrom(
		[]string{"csvtest"},
		[]string{"id", "bemerkung", "upd"},
		pgx.CopyFromRows(rower),
	); err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	} else {
		fmt.Println("nach copy", copyCount)
	}

	os.Exit(0)
}
