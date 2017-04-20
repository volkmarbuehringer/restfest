package db

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
)

var DBx *pgx.ConnPool

func InitDB() {

	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn
	//connConfig.Logger = log15.New("db", connConfig.Database)

	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 20, AfterConnect: setTyp}
	DBx, err = pgx.NewConnPool(config)
	if err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
	fmt.Println("stat", DBx.Stat())

}

//defer db.Close()
