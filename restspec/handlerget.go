package main

import (
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jackc/pgx"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

var decoder = schema.NewDecoder()

func getAllHandlerWeburl(w http.ResponseWriter, r *http.Request) {

	params := new(generrestspec.WeburlParams)

	err := r.ParseForm()

	if err != nil {
		service.SenderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	params.Length = 100
	err = decoder.Decode(params, r.Form)

	if err != nil {
		service.SenderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}
	var stmt *pgx.PreparedStatement
	if stmt, err = service.Prepare("weburl", service.GetSqlStmt(-2, 1), params); err != nil {

		return
	}
	rows, err := db.DBx.Query(stmt.Name, params.ROWInsert()...)
	if err != nil {
		service.SenderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}
	defer rows.Close()

	inter := make(generrestspec.ArWeburl, 0)
	err = inter.Scanner(rows)

	if err != nil {
		service.SenderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	service.Sender(w, inter)

}

func getByIDHandlerWeburl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	params := new(generrestspec.WeburlParams)
	if stmt, err := service.Prepare("weburl", service.GetSqlStmt(db.GenSelectID, 1), params); err != nil {
		service.SenderErr(w, err)
		log15.Error("DBFehler", "weburl", err)
		return
	} else {

		row := db.DBx.QueryRow(stmt.Name, id)

		weburl := new(generrestspec.Weburl)

		err = row.Scan(weburl.Scanner()...)
		if err != nil {
			service.SenderErr(w, err)
			log15.Error("DBFehler", "weburl", err)
			return
		} else {
			*weburl.Zusatz = 333 //set value in struct
			service.Sender(w, weburl)

		}

	}

}
