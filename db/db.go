package db

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/log15adapter"
	log "gopkg.in/inconshreveable/log15.v2"
)

var DBx *pgx.ConnPool

func InitDB() {

	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log.Crit("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn
	logger := log15adapter.NewLogger(log.New("module", "pgx"))
	connConfig.Logger = logger

	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 20, AfterConnect: SetTyp}
	DBx, err = pgx.NewConnPool(config)
	if err != nil {
		log.Crit("DB", "connect", err)
		os.Exit(1)
	}
	fmt.Println("stat", DBx.Stat())

}

//defer db.Close()
