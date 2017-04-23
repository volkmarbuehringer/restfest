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
	arrp := stru.Scanner()
	for anz := 0; rows.Next(); anz++ {

		if err = rows.Scan(arrp...); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return
		}
		if stru.Url != nil {
			fmt.Println(*stru.Url)
		}

		record, err := arrp.ConvertItoS()
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
	f1, err := os.Open("flat.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f1)

	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	iterator := copyCsv{
		rows:     r,
		structer: generteststruct.Weburl{},
	}

	iterator.inter = iterator.structer.Scanner()

	fmt.Println("vor copy", record)
	copyCount, err := dbx1.CopyFrom(
		[]string{"zielweburl"},
		iterator.structer.Columns(),
		&iterator)

	fmt.Println("fertig", copyCount, err)
	if err != nil {
		log.Fatal(err)
	}
}

type copyCsv struct {
	ptr      []string
	rows     *csv.Reader
	err      error
	inter    db.InterPgx
	structer generteststruct.Weburl
}

func (t *copyCsv) Next() bool {
	if t.err != nil {
		return false
	}
	for {
		t.ptr, t.err = t.rows.Read()

		if t.err != nil {
			return false
		}

		t.err = t.inter.ConvertStoI(t.ptr)
		if t.err != nil {
			return false
		}
		if t.structer.Url != nil {
			fmt.Println(*t.structer.Url)
			break
		} else {
			fmt.Println("niller", t.structer)

		}
	}

	return true
}
func (t *copyCsv) Values() ([]interface{}, error) {
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
	ptr   generteststruct.Weburl
	rows  *pgx.Rows
	inter db.InterPgx
}

func (t *WeburlScan) Next() bool {
	var ok bool
	for {
		ok = t.rows.Next()
		if !ok {
			break
		}
		t.rows.Scan(t.inter...)
		if t.ptr.Url != nil {
			fmt.Println(*t.ptr.Url)
			break
		}
	}
	return ok
}
func (t *WeburlScan) Values() ([]interface{}, error) {
	return t.inter, t.rows.Err()
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

	iterator := WeburlScan{
		ptr:  generteststruct.Weburl{},
		rows: rows,
	}
	iterator.inter = iterator.ptr.Scanner()
	fmt.Println("vor copy")
	copyCount, err := dbx1.CopyFrom(
		[]string{"zielweburl"},
		iterator.ptr.Columns(),
		&iterator)

	fmt.Println("fertig", copyCount, err)
}
func main() {

	csvread()

	copyer()
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
