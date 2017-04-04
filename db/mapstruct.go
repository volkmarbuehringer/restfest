package db

type SQLOper int

const (
	GenSelect SQLOper = iota
	GenInsert
	GenUpdate
)

type MapperFun1 func(interface{}) []interface{}

var SQLFunMap = map[string]func(string, SQLOper) []interface{}{}

var ROWInsertFunMap = map[string]MapperFun1{}

var EmptyFunMap = map[string]func() interface{}{}

var ScannerFunMap = map[string]func() ([]interface{}, interface{}){}
