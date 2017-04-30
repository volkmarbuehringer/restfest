package service

import (
	"fmt"
	"restfest/db"
	"strconv"

	"github.com/jackc/pgx"
)

var preparedStmt = map[string]*pgx.PreparedStatement{}

var ok bool

func GetSqlStmt(flagi db.SQLOper, tabFlag int) (flag db.SQLOper) {
	switch flagi {
	case -1:
		if tabFlag == 3 {
			flag = db.GenFunction
		} else {
			flag = db.GenInsert
		}
	case -2:
		switch tabFlag {
		case 3:
			flag = db.GenFunction
		case 4:
			flag = db.GenSelectAll1
		case 1, 2:
			flag = db.GenSelectAll
		default:
			//err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
			flag = flagi

		}
	default:
		flag = flagi
	}
	return
}

func PrepareSQL(sucher string, sql func() string) (stmt *pgx.PreparedStatement, err error) {

	if stmt, ok = preparedStmt[sucher]; !ok {

		if stmt, err = db.DBx.Prepare(sucher, sql()); err != nil {
			return
		} else {
			preparedStmt[sucher] = stmt
			fmt.Println(*stmt)
		}

	}
	return
}

func Prepare(tab string, flag db.SQLOper, params db.PgxGenerIns) (stmt *pgx.PreparedStatement, err error) {

	sucher := tab + strconv.Itoa(int(flag))
	if stmt, ok = preparedStmt[sucher]; !ok {
		sql := params.SQL(flag)
		if stmt, err = db.DBx.Prepare(sucher, sql); err != nil {
			return
		} else {
			preparedStmt[sucher] = stmt
			fmt.Println(*stmt)
		}

	}
	return
}

func ReadRow(tab string, id int, params db.PgxGenerIns) (inter db.PgxGener, err error) {

	if funMap, ok := db.FunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {
		if stmt, err1 := Prepare(tab, GetSqlStmt(db.GenSelectID, funMap.Flag), params); err1 != nil {

			return nil, err1
		} else {

			row := db.DBx.QueryRow(stmt.Name, id)

			inter = funMap.EmptyFun()

			err = row.Scan(inter.Scanner()...)

		}
	}

	return
}

func ReadRows(tab string, params db.PgxGenerIns) (rows *pgx.Rows, err error) {
	var stmt *pgx.PreparedStatement
	if funMap, ok := db.FunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {
		if stmt, err = Prepare(tab, GetSqlStmt(-2, funMap.Flag), params); err != nil {

			return
		}
		rows, err = db.DBx.Query(stmt.Name, params.ROWInsert()...)

	}
	return
}
