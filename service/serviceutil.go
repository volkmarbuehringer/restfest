package service

import (
	"fmt"
	"restfest/db"
	"strconv"

	"github.com/jackc/pgx"
)

var preparedStmt = map[string]*pgx.PreparedStatement{}

func row1Scanner(funMap db.TFunMap, rows *pgx.Row) (stru db.PgxGener, err error) {
	stru = funMap.EmptyFun()

	arr := stru.Scanner()
	if err = rows.Scan(arr...); err != nil {

		return
	}
	return

}

func rowScanner(funMap db.TFunMap, rows *pgx.Rows) (stru []db.PgxGener, err error) {
	t := make([]db.PgxGener, 0)

	var x db.PgxGener

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return
		}
		x = funMap.EmptyFun()
		arr := x.Scanner()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		t = append(t, x)
	}
	stru = t
	return
}

func prepare(tab string, flag db.SQLOper) (stmt *pgx.PreparedStatement, sqlFun db.TFunMap, err error) {
	var ok bool

	if sqlFun, ok = db.FunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	}

	if flag < 0 {
		if sqlFun.Flag == 3 {
			flag = db.GenFunction
		} else {
			flag = db.GenInsert
		}

	}

	sucher := tab + strconv.Itoa(int(flag))
	if stmt, ok = preparedStmt[sucher]; !ok {
		if stmt, err = db.DBx.Prepare(sucher, sqlFun.SQLFun(flag)); err != nil {
			return
		} else {
			preparedStmt[sucher] = stmt
			fmt.Println(*stmt)
		}

	}
	return
}

func readRow(tab string, id int) (inter db.PgxGener, err error) {
	if stmt, funMap, err1 := prepare(tab, db.GenSelectID); err1 != nil {

		return nil, err1
	} else {

		rows := db.DBx.QueryRow(stmt.Name, id)

		inter, err = row1Scanner(funMap, rows)

	}
	return
}

func readRows(tab string, sqler db.SQLOper, params db.PgxGenerIns) (inter []db.PgxGener, err error) {
	var stmt *pgx.PreparedStatement
	var sqlFun db.TFunMap
	if stmt, sqlFun, err = prepare(tab, sqler); err != nil {

		return
	}
	rows, err := db.DBx.Query(stmt.Name, params.ROWInsert()...)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	inter, err = rowScanner(sqlFun, rows)
	return

}
