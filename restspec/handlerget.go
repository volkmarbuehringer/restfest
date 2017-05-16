package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/restspec/gener"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
	"github.com/jackc/pgx"
)

func getAllHandlerLos(w http.ResponseWriter, r *http.Request) {

	var params = gener.LosParams{
		Length: 100, //default read 100 rows
	}

	err := service.PrepParam(&params, w, r)
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	var stmt *pgx.PreparedStatement

	if stmt, err = service.PrepareSQL("losAll", func() string {
		return fmt.Sprintf("select %s from "+db.DBschema+".los where l_iban is null limit $1 offset $2",
			gener.LosSQL.All,
		)
	}); err != nil {
		service.SenderErr(w, r, err)
		return
	}
	rows, err := db.DBx.Query(stmt.Name, params.ROWInsert()...)
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	defer rows.Close()

	var iter gener.IterLos
	iter.NewCopy(rows) //streaming from database

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("["))
	var trenner string
	for anz := 0; iter.Next(); anz++ {
		iter.Los.L_iban = db.String("efsdfsadfsdf")
		if iter.Los.L_iban != nil {

			if anz > 0 {
				trenner = ","
			}
			err = iter.Los.Writer(w, trenner)
			if err != nil {
				service.SenderErr(w, r, err)
				return
			}

		}

	}
	if iter.Errc != nil {
		service.SenderErr(w, r, iter.Errc)
		return
	}
	w.Write([]byte("]"))

}

func getByIDHandlerLos(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	if stmt, err := service.PrepareSQL("losID", func() string {
		return fmt.Sprintf("select %s from "+db.DBschema+".los where l_id = $1 ",
			gener.LosSQL.All,
		)
	}); err != nil {
		service.SenderErr(w, r, err)
		return
	} else {

		row := db.DBx.QueryRow(stmt.Name, id)

		los := new(gener.Los)

		err = row.Scan(los.Scanner()...)
		if err != nil {
			service.SenderErr(w, r, err)
			return
		} else {
			los.L_iban = db.String("neuer wert")

			service.Sender(w, los)

		}

	}

}
