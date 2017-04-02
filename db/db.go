package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Pager struct {
	Length int
	Offset int
	ID     int
	Where  string
}

func init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable application_name=gogogo connect_timeout=3",
		os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	DB, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(1000 * time.Second)
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("stats", DB.Stats())

}

//defer db.Close()
