package db

import "os"

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

type MapperFun1 func(interface{}) []interface{}

var SQLFunMap = map[string]func(SQLOper) string{}

var ROWInsertFunMap = map[string]MapperFun1{}

var ROWQueryFunMap = map[string]MapperFun1{}

var EmptyFunMap = map[string]func() interface{}{}

var ParamFunMap = map[string]func() interface{}{}

var ScannerFunMap = map[string]func() ([]interface{}, interface{}){}

var FlagMap = map[string]int{}
