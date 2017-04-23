package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type SQLOper int

type PgxGenerAr interface {
	Scanner(*pgx.Rows) error
}
type PgxGenerMap interface {
	Scanner(*pgx.Rows) error
}
type PgxGener interface {
	Scanner() InterPgx
	ROWInsert() InterPgx
}

type PgxGenerIns interface {
	ROWInsert() InterPgx
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

func (arr InterPgx) ConvertItoS() (record []string, err error) {
	record = make([]string, len(arr))

	for i := range arr {
		switch arr[i].(type) {
		case string:
			record[i] = arr[i].(string)
		case *string:
			t := arr[i].(*string)
			if t != nil {
				record[i] = *t
			} else {
				record[i] = ""
			}
		case int32:
			record[i] = strconv.Itoa((int)(arr[i].(int32)))
		case int64:
			record[i] = strconv.Itoa((int)(arr[i].(int64)))
		case *int32:
			t := arr[i].(*int32)
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case *int64:
			t := arr[i].(*int64)
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case float64:
			record[i] = strconv.FormatFloat(arr[i].(float64), 'E', -1, 64)
		case *float64:
			t := arr[i].(*float64)
			if t != nil {
				record[i] = strconv.FormatFloat(*t, 'E', -1, 64)
			} else {
				record[i] = ""
			}
		case bool:
			record[i] = strconv.FormatBool(arr[i].(bool))
		case *time.Time:
			t := arr[i].(*time.Time)
			if t != nil {
				record[i] = t.Format("2006-01-01")
			} else {
				record[i] = ""
			}
		case time.Time:
			record[i] = (arr[i].(time.Time)).Format("2006-01-01")
		default:
			if arr[i] != nil {
				err = fmt.Errorf("unknown %d %T", i, arr[i])

			} else {
				err = fmt.Errorf("unknown %d ", i)

			}

		}

	}
	return
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
	SQLFun      func(SQLOper) string
	EmptyFun    func() PgxGener
	EmptyInsFun func() PgxGenerIns
	ParamFun    func() PgxGenerIns
	EmptyArray  func() PgxGenerAr
	Flag        int
}

var FunMap = map[string]TFunMap{}

var ConnectorFunMap = map[string]MapperFun2{}
