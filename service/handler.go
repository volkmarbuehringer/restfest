package service

import (
	"net/http"
	"restfest/db"
	"strconv"

	"github.com/gorilla/mux"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

func getByIDHandler3(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]

	ga, err := prepParam(tab, w, r)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	var sqler db.SQLOper
	if db.FlagMap[tab] == 3 {
		sqler = db.GenFunction
	} else {
		sqler = db.GenSelectAll
	}
	if stmt, err1 := prepare(tab, sqler); err1 != nil {
		senderErr(w, err1)
		log15.Error("DBFehler", "getall", err1)
		return
	} else {
		rows, err := db.DBx.Query(stmt.Name, db.ROWQueryFunMap[tab](ga)...)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}
		defer rows.Close()

		inter, err := rowScanner(tab, rows)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}

		sender(w, inter)

	}

}

func getByIDHandler5(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	if inter, err := readRow(tab, id); err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	} else {
		sender(w, inter)
	}

}
