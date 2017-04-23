package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
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

	f1, err := os.Create("flat.csv")
	if err != nil {
		log15.Crit("DBFehler", "create", err)
		log.Fatal(err)
	}
	w := csv.NewWriter(f1)

	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer rows.Close()
	stru := new(generteststruct.Weburl)
	for anz := 0; rows.Next(); anz++ {
		arrp := stru.Scanner()
		if err = rows.Scan(arrp...); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return
		}

		record, err := stru.ScannerV().ConvertItoS()
		if err != nil {
			log.Fatal(err)
		}

		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
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

func csvread() {
	f1, err := os.Open("/home/walter/Downloads/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	/*
		r := csv.NewReader(f1)

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(record)
		}
	*/
	r := csv.NewReader(f1)

	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	x := new(generteststruct.Csvtest)
	iterator := copyCsv{
		rows:  r,
		inter: x.ScannerVI(),
	}

	fmt.Println("vor copy", record)
	copyCount, err := dbx1.CopyFrom(
		[]string{"csvtest"},
		x.Columns(),
		&iterator)

	fmt.Println("fertig", copyCount, err)
	if err != nil {
		log.Fatal(err)
	}
}

type copyCsv struct {
	ptr   []string
	rows  *csv.Reader
	err   error
	inter db.InterPgx
}

func (t *copyCsv) Next() bool {
	t.ptr, t.err = t.rows.Read()

	if t.err != nil {
		return false
	}
	return true
}
func (t *copyCsv) Values() ([]interface{}, error) {

	t.err = t.inter.ConvertStoI(t.ptr)

	return t.inter, t.err
}
func (t *copyCsv) Err() error {
	if t.err != io.EOF {
		return t.err
	} else {
		return nil
	}
}

type WeburlScan struct {
	ptr   *generteststruct.Weburl
	rows  *pgx.Rows
	inter db.InterPgx
}

func (t *WeburlScan) Next() bool {
	return t.rows.Next()
}
func (t *WeburlScan) Values() ([]interface{}, error) {
	err := t.rows.Scan(t.inter...)
	return t.inter, err
}
func (t *WeburlScan) Err() error {
	return t.rows.Err()
}
func copyer() {
	rows, err := dbx.Query(generteststruct.SQLWeburl(db.GenSelectAll1), 10000, 0)
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer rows.Close()

	x := new(generteststruct.Weburl)
	iterator := WeburlScan{
		x,
		rows,
		x.Scanner(),
	}

	fmt.Println("vor copy")
	copyCount, err := dbx1.CopyFrom(
		[]string{"zielweburl"},
		x.Columns(),
		&iterator)

	fmt.Println("fertig", copyCount, err)
}
func main() {

	csvread()

	//copyer()
	//mapper()
	//fetcher()
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
