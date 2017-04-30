package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

func getAllHandlerLos(w http.ResponseWriter, r *http.Request) {

	var params = generrestspec.LosParams{
		Length: 100, //default read 100 rows
	}

	err := service.PrepParam(&params, w, r)
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	var stmt *pgx.PreparedStatement

	if stmt, err = service.PrepareSQL("losAll", func() string {
		return fmt.Sprintf("select %s from "+db.DBschema+".los where l_iban is not null limit $1 offset $2",
			generrestspec.LosSQL.All,
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

	var iter generrestspec.IterLos
	iter.NewCopy(rows) //streaming from database

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("["))
	for anz := 0; iter.Next(); anz++ {

		if iter.Errc != nil {
			service.SenderErr(w, iter.Errc)
			return
		}
		if iter.Los.L_iban != nil {
			*iter.Los.L_iban = "efsdfsadfsdf"
			if anz > 0 {
				w.Write([]byte(","))
			}
			err = iter.Los.Writer(w)
			if err != nil {
				service.SenderErr(w, err)
				return
			}

		}

	}
	w.Write([]byte("]"))

}

func getByIDHandlerLos(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if stmt, err := service.PrepareSQL("losID", func() string {
		return fmt.Sprintf("select %s from "+db.DBschema+".los where id = $1 ",
			generrestspec.LosSQL.All,
		)
	}); err != nil {
		service.SenderErr(w, err)
		return
	} else {

		row := db.DBx.QueryRow(stmt.Name, id)

		los := new(generrestspec.Los)

		err = row.Scan(los.Scanner()...)
		if err != nil {
			service.SenderErr(w, err)
			return
		} else {
			if los.L_iban != nil {
				*los.L_iban = "333" //set value in struct
			}

			service.Sender(w, los)

		}

	}

}
