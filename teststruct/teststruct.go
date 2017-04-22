package main

import (
	"fmt"
	"os"
	"restfest/db"
	"restfest/generteststruct"

	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

//go:generate go run ../gen.go teststruct 1

var dbx *pgx.Conn

func main() {

	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	mapper := make(generteststruct.MapWeburl)
	mapper.Scanner(rows)

	for r, m := range mapper {
		if m.Url != nil {
			fmt.Println(r, *m.Url)
		}

	}
	defer rows.Close()
	/*
		stru := new(generteststruct.Weburl)
		for anz := 0; rows.Next(); anz++ {

			if err = rows.Scan(stru.Scanner()...); err != nil {
				log15.Crit("DBFehler", "scan", err)
				return
			}
			fmt.Println(anz, stru)

		}
	*/
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

	err = db.SetTyp(dbx)
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
}
