package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

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
	DBx, err = pgx.NewConnPool(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("stat", DBx.Stat())

}

//defer db.Close()
