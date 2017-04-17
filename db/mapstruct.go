package db

import (
	"os"

	"github.com/jackc/pgx"
)

type SQLOper int

const (
	GenSelectID SQLOper = iota
	GenInsert
	GenUpdate
	GenFunction
	GenSelectAll
	GenDelete
	GenSelectAll1
)

var DBschema = os.Getenv("PGSCHEMA")

var SQLPattern = []string{"select %s from " + DBschema + ".%s where %s=$1",
	`insert into ` + DBschema + `.%s(%s)values(%s) returning %s`,
	`update ` + DBschema + `.%s set %s where %s returning %s`,
	"select %s from " + DBschema + ".%s ( %s )",
	"select %s from " + DBschema + ".%s order by %s limit $1 offset $2",
	`delete from ` + DBschema + `.%s where %s=$1 returning %s`,
	"select %s from " + DBschema + ".%s",
}

func setTyp(con *pgx.Conn) error {
	for _, f := range ConnectorFunMap {
		f(con)
	}
	return nil
}

type MapperFun1 func(interface{}) InterPgx

var SQLFunMap = map[string]func(SQLOper) string{}

var ROWInsertFunMap = map[string]MapperFun1{}

var ROWQueryFunMap = map[string]MapperFun1{}

var EmptyFunMap = map[string]func() interface{}{}

var ParamFunMap = map[string]func() interface{}{}

var ScannerFunMap = map[string]func() (InterPgx, interface{}){}

var ConnectorFunMap = map[string]func(*pgx.Conn) error{}

var FlagMap = map[string]int{}
