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
var dbx1 *pgx.Conn

func fetcher() {
	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer rows.Close()
	stru := new(generteststruct.Weburl)
	for anz := 0; rows.Next(); anz++ {

		if err = rows.Scan(stru.Scanner()...); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return
		}
		fmt.Println(anz, stru)

	}

}
func mapper() {
	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer rows.Close()
	mapper := make(generteststruct.MapWeburl)
	mapper.Scanner(rows)

	for r, m := range mapper {
		if m.Url != nil {
			fmt.Println(r, *m.Url)
		}

	}

}
func main() {

	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer rows.Close()
	iterator := generteststruct.WeburlIterator(rows)

	fmt.Println("vor copy")
	copyCount, err := dbx1.CopyFrom(
		[]string{"zielweburl"},
		iterator.Columns(),
		iterator)

	fmt.Println("fertig", copyCount, err)
	mapper()
	fetcher()
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

	if dbx, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	//	fmt.Println(dbx.ConnInfo)
	if dbx1, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	err = db.SetTyp(dbx)
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
}
