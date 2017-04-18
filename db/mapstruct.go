package db

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
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
	"select %s from " + DBschema + ".%s limit $1 offset $2",
}

type oidmerker struct {
	oida pgtype.Oid
	oid  pgtype.Oid
}

var dbschema = os.Getenv("PGSCHEMA")
var mapper = map[string]oidmerker{}

func setTyp(con *pgx.Conn) error {

	if len(mapper) == 0 {
		rower, err := con.Query(`select t.oid, t.typname,t.typarray
	from pg_type t
	where (
		  t.typtype   not in('b', 'p', 'r')
	    and t.typarray > 0
		)
	    and t.typnamespace = ( select oid from pg_namespace where nspname = '` + dbschema + `')
	`)
		if err != nil {
			return err
		}

		mapper = make(map[string]oidmerker)

		defer rower.Close()
		for rower.Next() {
			var oida, oid pgtype.Oid
			var name string
			err := rower.Scan(&oid, &name, &oida)
			if err != nil {
				return err
			}
			mapper[name] = oidmerker{oida, oid}
			fmt.Println("reigster", name)
		}

	}

	for r, f := range ConnectorFunMap {
		if fz, ok := mapper[r]; !ok {
			return fmt.Errorf("tab nicht gefunden %s", r)
		} else {

			f(con, fz.oida, fz.oid)

		}

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

var ConnectorFunMap = map[string]func(*pgx.Conn, pgtype.Oid, pgtype.Oid) error{}

var FlagMap = map[string]int{}
