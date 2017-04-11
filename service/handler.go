package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"restfest/db"
)

func getByIDHandler3(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]

	ga, err := prepParam(tab, w, r)
	if err != nil {
		senderErr(w, err)
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
		return
	} else {
		rows, err := db.DBx.Query(stmt.Name, db.ROWQueryFunMap[tab](ga)...)
		if err != nil {
			senderErr(w, err)
			return
		}
		defer rows.Close()

		inter, err := rowScanner(tab, rows, 0)
		if err != nil {
			senderErr(w, err)
			return
		}

		sender(w, inter)

	}

}

func getByIDHandler5(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	if stmt, err1 := prepare(tab, db.GenSelectID); err1 != nil {
		senderErr(w, err1)
		return
	} else {
		rows, err := db.DBx.Query(stmt.Name, id)
		if err != nil {
			senderErr(w, err)
			return
		}
		defer rows.Close()

		if inter, err := rowScanner(tab, rows, 1); err != nil {
			senderErr(w, err)
			return
		} else {
			sender(w, inter)
		}

	}

}
