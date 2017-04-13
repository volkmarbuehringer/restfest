package main

import (
	"fmt"
	"os"
	"restfest/db"
	"restfest/gener"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

var dbx *pgx.Conn

type Copy struct {
	rows *pgx.Rows
}

func (t Copy) Next() bool {
	if anz >= limiter {
		return false
	}
	anz++

	return t.rows.Next()
}

func (t Copy) Values() (arr []interface{}, err error) {
	arrr, _ := gener.ScannerTZahlung()
	if err = t.rows.Scan(arrr...); err != nil {
		log15.Crit("DBFehler", "scan", err)
		return
	}
	return arrr, nil
}

func (t Copy) Err() error {
	return t.rows.Err()
}

var limiter int
var anz int

func main() {

	limiter, _ = strconv.Atoi(os.Args[1])
	var dat Copy
	var err error
	dat.rows, err = dbx.Query(gener.SQLZahlung(db.GenSelectAll1))
	if err != nil {
		log15.Crit("DBFehler", "get", err)
		return
	}
	defer dat.rows.Close()

	var datI pgx.CopyFromSource

	datI = dat

	//store := make(map[int64]*gener.Zahlung, 0)
	fmt.Println("vor copy")
	t := time.Now()

	copyCount, err := db.DBx.CopyFrom(
		[]string{"copyzahlung"},
		[]string{"z_id", "z_jobid", "z_packageid", "z_ticketcompid", "z_paytypeid", "z_wsid", "z_ezid", "z_flid", "z_hznr_buchung", "z_betrag", "z_zahlungsdatum", "z_blz", "z_konto", "z_bankref", "z_blzempf", "z_kontoempf", "z_loeschkenn", "z_cr_uid", "z_cr_date", "z_upd_uid", "z_upd_date", "z_bicempf", "z_ibanempf", "z_bic", "z_iban"},
		datI,
	)

	if err != nil {
		log15.Crit("DBFehler", "copy", err)
		return
	}

	//		store[struT.Z_id] = struT
	fmt.Println("nach copy", copyCount, time.Since(t))
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
}
