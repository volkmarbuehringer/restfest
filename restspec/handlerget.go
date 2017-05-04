package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	gener "restfest/restspec/gener"
	"restfest/service"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/julienschmidt/httprouter"
)

func getAllHandlerLos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var params = gener.LosParams{
		Length: 100, //default read 100 rows
	}

	err := service.PrepParam(&params, w, r)
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	var stmt *pgx.PreparedStatement

	if stmt, err = service.PrepareSQL("losAll", func() string {
		return fmt.Sprintf("select %s from "+db.DBschema+".los where l_iban is null limit $1 offset $2",
			gener.LosSQL.All,
		)
	}); err != nil {
		service.SenderErr(w, err)
		return
	}
	rows, err := db.DBx.Query(stmt.Name, params.ROWInsert()...)
	if err != nil {
		service.SenderErr(w, err)
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
				service.SenderErr(w, err)
				return
			}

		}

	}
	if iter.Errc != nil {
		service.SenderErr(w, iter.Errc)
		return
	}
	w.Write([]byte("]"))

}

func getByIDHandlerLos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	if stmt, err := service.PrepareSQL("losID", func() string {
		return fmt.Sprintf("select %s from "+db.DBschema+".los where id = $1 ",
			gener.LosSQL.All,
		)
	}); err != nil {
		service.SenderErr(w, err)
		return
	} else {

		row := db.DBx.QueryRow(stmt.Name, id)

		los := new(gener.Los)

		err = row.Scan(los.Scanner()...)
		if err != nil {
			service.SenderErr(w, err)
			return
		} else {
			los.L_iban = db.String("neuer wert")

			service.Sender(w, los)

		}

	}

}
