package db

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"

	"github.com/sirupsen/logrus"
)

var DBx *pgx.ConnPool

func InitDB() {

	log := logrus.New()
	log.Out = os.Stdout
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log.Fatal("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn
	logger := logrusadapter.NewLogger(log)
	connConfig.Logger = logger

	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 20, AfterConnect: SetTyp}
	DBx, err = pgx.NewConnPool(config)
	if err != nil {
		log.Fatal("DB", "connect", err)
		os.Exit(1)
	}
	fmt.Println("stat", DBx.Stat())

}

//defer db.Close()
