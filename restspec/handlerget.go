package main

import (
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

func getAllHandlerWeburl(w http.ResponseWriter, r *http.Request) {

	var params = generrestspec.WeburlParams{
		Length: 100, //default read 100 rows
	}

	err := service.PrepParam(&params, w, r)
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	var stmt *pgx.PreparedStatement
	if stmt, err = service.Prepare("weburl", db.GenSelectAll, &params); err != nil {
		service.SenderErr(w, err)
		return
	}
	rows, err := db.DBx.Query(stmt.Name, params.ROWInsert()...)
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	defer rows.Close()

	var iter generrestspec.IterWeburl
	iter.NewCopy(rows) //streaming from database

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("["))
	for anz := 0; iter.Next(); anz++ {

		if iter.Errc != nil {
			service.SenderErr(w, iter.Errc)
			return
		}
		if iter.Weburl.Zusatz != nil {
			*iter.Weburl.Zusatz = 7899
			if anz > 0 {
				w.Write([]byte(","))
			}
			err = iter.Weburl.Writer(w)
			if err != nil {
				service.SenderErr(w, err)
				return
			}

		}

	}
	w.Write([]byte("]"))

}

func getByIDHandlerWeburl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	params := new(generrestspec.WeburlParams)
	if stmt, err := service.Prepare("weburl", db.GenSelectID, params); err != nil {
		service.SenderErr(w, err)
		return
	} else {

		row := db.DBx.QueryRow(stmt.Name, id)

		weburl := new(generrestspec.Weburl)

		err = row.Scan(weburl.Scanner()...)
		if err != nil {
			service.SenderErr(w, err)
			return
		} else {
			if weburl.Zusatz != nil {
				*weburl.Zusatz = 333 //set value in struct
			}

			service.Sender(w, weburl)

		}

	}

}
