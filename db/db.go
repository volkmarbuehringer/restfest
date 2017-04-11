package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

var DB *sql.DB
var DBx *pgx.ConnPool

func init() {

	porter, err := strconv.Atoi(os.Getenv("PGPORT"))
	connConfig := pgx.ConnConfig{
		Host:     os.Getenv("PGHOST"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		Database: os.Getenv("PGDATABASE"),
		Port:     uint16(porter),
	}

	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 20}
	pool, err := pgx.NewConnPool(config)
	if err != nil {
		log.Fatal(err)
	}

	DB, err = stdlib.OpenFromConnPool(pool)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	if driver, ok := DB.Driver().(*stdlib.Driver); ok && driver.Pool != nil {
		// fast path with pgx
		DBx = driver.Pool
		fmt.Println("stat", DBx.Stat())
	} else {
		log.Fatal("alles kaputt")
	}

}

//defer db.Close()
