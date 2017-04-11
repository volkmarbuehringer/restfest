package db

import (
	"fmt"

	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
)

var DBx *pgx.ConnPool

func init() {

	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log15.Crit("DBFehler", err)
	}
	connConfig.LogLevel = pgx.LogLevelWarn

	connConfig.Logger = log15.New("db", connConfig.Database)
	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 20}
	DBx, err = pgx.NewConnPool(config)
	if err != nil {
		log15.Crit("DBFehler", err)
	}
	fmt.Println("stat", DBx.Stat())

}

//defer db.Close()
