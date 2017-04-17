package main

import (
	"fmt"
	"os"
	"restfest/db"
	"restfest/generteststruct"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

//go:generate go run ../gen.go teststruct 1

var dbx *pgx.Conn

func main() {

	rows, err := dbx.Query(generteststruct.SQLGuges1(db.GenSelectAll1))
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer rows.Close()

	for anz := 0; rows.Next(); anz++ {
		stru := new(generteststruct.Guges1)
		arrr := generteststruct.ScannerGuges1I(stru)
		if err = rows.Scan(arrr...); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return
		}
		fmt.Println("guer da", anz, len(stru.Agger), *stru.Texter, *stru.Texter2, *stru.Zahler, *stru.Zahler2)
		for i, w := range stru.Agger {
			if w.Url != nil {
				println(i, *w.Url)
			}
		}

	}
	defer dbx.Close()

	os.Exit(0)
}

func init() {
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn

	var dbe pgx.Conn
	if dbx, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	fmt.Println(dbe.ConnInfo)
	//fmt.Println(gaga.ConnInfo)
	dbe.ConnInfo.RegisterDataType(pgtype.DataType{
		Value: &generteststruct.Weburls{},
		Name:  "weburls",
		Oid:   24877,
	})

}
