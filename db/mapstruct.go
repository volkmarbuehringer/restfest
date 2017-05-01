package db

import (
	"fmt"
	"io"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type SQLOper int

type Iterator interface {
	Next() bool
	NewCopy(rows *pgx.Rows)
	Err() error
	Value() PgxGener
}
type PgxGenerMap interface {
	Scanner(*pgx.Rows) error
}
type PgxGener interface {
	Scanner() InterPgx
	ROWInsert() InterPgx
	Columns() []string
	SQL(SQLOper) string
	Reader(map[string][]string) [][]string
	Writer(io.Writer, string) error
}

type PgxGenerIns interface {
	ROWInsert() InterPgx
	SQL(SQLOper) string

	Scanner() InterPgx
	Reader(map[string][]string) [][]string
}

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

func LoadTypMap(con *pgx.Conn) error {

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
	return nil
}

func init() {

}

func SetTyp(con *pgx.Conn) error {

	if err := LoadTypMap(con); err != nil {
		return err
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

type MapperFun2 func(*pgx.Conn, pgtype.Oid, pgtype.Oid) error

type TFunMap struct {
	EmptyFun    func() PgxGener
	EmptyInsFun func() PgxGenerIns
	ParamFun    func() PgxGenerIns
	Iterator    func() Iterator
	Flag        int
}

var FunMap = map[string]TFunMap{}

var ConnectorFunMap = map[string]MapperFun2{}
