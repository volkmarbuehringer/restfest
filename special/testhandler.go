package main

import (
	"fmt"
	"log"
	"os"
	"restfest/db"
	"restfest/gener"
	"strconv"

	"github.com/jackc/pgx"
)

func main() {

	limiter, _ := strconv.Atoi(os.Args[1])
	rows, err := db.DBx.Query(fmt.Sprintf("select %s from gutschein", gener.SQLGutschein(db.GenSelect)[0]))
	if err != nil {
		return
	}
	store := make(map[int64]*gener.Gutschein, 0)
	defer rows.Close()
	for anz := 0; rows.Next(); anz++ {
		arr, ts := gener.ScannerTGutschein()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		store[ts.Aug_id.Int64] = ts
		if ts.Auf_anzahl.Int64 > 1 {
			fmt.Printf("grösser %s %s\n", ts.Auf_upd_uid, ts.Auf_bemerkung)
		}

		if len(ts.Auf_partner.String) > 0 {
			fmt.Printf("%d %s \n", ts.Auf_id.Int64, ts.Auf_partner)
		} else {
			fmt.Printf("%d null \n", ts.Auf_id.Int64)
		}
		//fmt.Printf("%v %T \n", ts, ts)
		if anz > limiter {
			break
		}
	}
	defer db.DBx.Close()

	rower := make([][]interface{}, 0)
	for _, t := range store {
		rower = append(rower, []interface{}{t.Aug_id, t.Auf_bemerkung, t.Auf_upd_uid})
	}
	fmt.Println("vor copy", len(rower))
	copyCount, err1 := db.DBx.CopyFrom(
		[]string{"csvtest"},
		[]string{"id", "bemerkung", "upd"},
		pgx.CopyFromRows(rower),
	)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("nach copy", copyCount)
	/*

		txn, err := db.DBx.Begin()

		if err != nil {
			log.Fatal(err)
		}
			stmt, err := txn.Prepare(pq.CopyIn("csvtest", "id", "bemerkung", "upd"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("starte copy ", len(store))
			for _, t := range store {
				if _, err = stmt.Exec(t.Aug_id, t.Auf_bemerkung, t.Auf_upd_uid); err != nil {
					log.Fatal(err)
				}
			}
			_, err = stmt.Exec()
			if err != nil {
				log.Fatal(err)
			}

			err = stmt.Close()
			if err != nil {
				log.Fatal(err)
			}

			err = txn.Commit()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("nach copy ", len(store))
	*/
	os.Exit(0)
}
