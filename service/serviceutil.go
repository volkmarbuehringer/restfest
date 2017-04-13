package service

import (
	"fmt"
	"restfest/db"
	"strconv"

	"github.com/jackc/pgx"
)

var preparedStmt = map[string]*pgx.PreparedStatement{}

func row1Scanner(tab string, rows *pgx.Row) (stru interface{}, err error) {
	var arr []interface{}
	arr, stru = db.ScannerFunMap[tab]()
	if err = rows.Scan(arr...); err != nil {

		return
	}
	return

}

func rowScanner(tab string, rows *pgx.Rows) (stru interface{}, err error) {
	t := make([]interface{}, 0)
	fun := db.ScannerFunMap[tab]

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return
		}
		arr, ts := fun()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		t = append(t, ts)
	}
	stru = &t
	return
}

func prepare(tab string, flag db.SQLOper) (stmt *pgx.PreparedStatement, err error) {
	sucher := tab + strconv.Itoa(int(flag))
	var ok bool
	if stmt, ok = preparedStmt[sucher]; !ok {
		if sqlFun, ok := db.SQLFunMap[tab]; !ok {
			err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
			return
		} else if stmt, err = db.DBx.Prepare(sucher, sqlFun(flag)); err != nil {
			return
		} else {
			preparedStmt[sucher] = stmt
			fmt.Println(*stmt)
		}

	}
	return
}

func readRow(tab string, id int) (inter interface{}, err error) {
	if stmt, err1 := prepare(tab, db.GenSelectID); err1 != nil {

		return nil, err1
	} else {

		rows := db.DBx.QueryRow(stmt.Name, id)

		inter, err = row1Scanner(tab, rows)

	}
	return
}
