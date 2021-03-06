package main

import (
	"fmt"
	"log"
	"os"
	"restfest/db"
	"restfest/teststruct/gener"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jackc/pgx"
)

//go:generate go run ../gen.go teststruct 1

var dbx *pgx.Conn
var dbx1 *pgx.Conn
var dbx2 *pgx.Conn
var dbx3 *pgx.Conn
var dbx4 *pgx.Conn

func mapper() error {
	params := new(gener.LosParams)
	rows, err := dbx4.Query(params.SQL(db.GenSelectAll1), 5000000, 0)
	if err != nil {
		logrus.Fatal("DBFehler", "get", err)
		return err
	}
	defer rows.Close()
	var iter gener.MapIterLos
	anz := iter.NewCopy(rows) //streaming from database

	fmt.Println("len of map", anz)
	/*
		for r, m := range mapper {
			if m.L_iban != nil {
				fmt.Println(r, *m.L_iban)
			}

		}
	*/
	return rows.Err()

}

func runner(funcer func() error, c chan error, te string) {
	start := time.Now()
	err := funcer()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(te, time.Since(start))
	c <- err
}

func main() {
	defer dbx.Close()
	defer dbx1.Close()
	defer dbx2.Close()

	c := make(chan error)
	go runner(mapper, c, "mapper")
	go runner(fetcher, c, "fetcher")
	go runner(csvread, c, "csvread")
	go runner(copyer, c, "copy table")

	fmt.Println("run", <-c)
	fmt.Println("run", <-c)
	fmt.Println("run", <-c)
	fmt.Println("run", <-c)

	os.Exit(0)
}

func init() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn

	if dbx, err = pgx.Connect(connConfig); err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}
	//	fmt.Println(dbx.ConnInfo)
	if dbx1, err = pgx.Connect(connConfig); err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}
	if dbx2, err = pgx.Connect(connConfig); err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}
	if dbx3, err = pgx.Connect(connConfig); err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}
	if dbx4, err = pgx.Connect(connConfig); err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}

	err = db.SetTyp(dbx)
	if err != nil {
		logrus.Fatal("DB", "parse", err)
		os.Exit(1)
	}
}
