package db

type SQLOper int

const (
	GenSelectID SQLOper = iota
	GenInsert
	GenUpdate
	GenFunction
	GenSelectAll
	GenDelete
)

type MapperFun1 func(interface{}) []interface{}

var SQLFunMap = map[string]func(SQLOper) string{}

var ROWInsertFunMap = map[string]MapperFun1{}

var ROWQueryFunMap = map[string]MapperFun1{}

var EmptyFunMap = map[string]func() interface{}{}

var ParamFunMap = map[string]func() interface{}{}

var ScannerFunMap = map[string]func() ([]interface{}, interface{}){}

var FlagMap = map[string]int{}
