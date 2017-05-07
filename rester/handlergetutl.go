package main

import (
	"fmt"
	"restfest/db"
	"restfest/service"

	"github.com/jackc/pgx"
)

func ReadRow(tab string, id int, params db.PgxGenerIns) (inter db.PgxGener, err error) {

	if funMap, ok := db.FunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {
		if stmt, err1 := service.Prepare(tab, service.GetSqlStmt(db.GenSelectID, funMap.Flag), params); err1 != nil {

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
		if stmt, err = service.Prepare(tab, service.GetSqlStmt(-2, funMap.Flag), params); err != nil {

			return
		}
		rows, err = db.DBx.Query(stmt.Name, params.ROWInsert()...)

	}
	return
}
