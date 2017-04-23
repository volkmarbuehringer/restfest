package main

import (
	"fmt"
	"log"
	"os"
	"restfest/db"
	"restfest/generteststruct"
	"time"

	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

//go:generate go run ../gen.go teststruct 1

var dbx *pgx.Conn
var dbx1 *pgx.Conn
var dbx2 *pgx.Conn

func mapper() error {
	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return err
	}
	defer rows.Close()
	mapper := make(generteststruct.MapWeburl)
	mapper.Scanner(rows)

	for r, m := range mapper {
		if m.Url != nil {
			fmt.Println(r, *m.Url)
		}

	}
	return rows.Err()

}

func main() {

	defer dbx.Close()
	defer dbx1.Close()
	defer dbx2.Close()

	start := time.Now()

	if err := csvread(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("copy flatfile", time.Since(start))

	start = time.Now()
	if err := mapper(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("map table", time.Since(start))
	start = time.Now()
	if err := fetcher(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("read flatfile", time.Since(start))

	start = time.Now()

	if err := copyer(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("copy table", time.Since(start))
	os.Exit(0)
}

func init() {
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn

	if dbx, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	//	fmt.Println(dbx.ConnInfo)
	if dbx1, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	if dbx2, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	err = db.SetTyp(dbx)
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
}
