package service

import (
	"net/http"
	"restfest/db"
	"strconv"

	"github.com/gorilla/mux"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]

	params, err := prepParam(tab, w, r)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	var sqler db.SQLOper
	switch db.FlagMap[tab] {
	case 3:
		sqler = db.GenFunction
	case 4:
		sqler = db.GenSelectAll1
	default:
		sqler = db.GenSelectAll
	}
	inter, err := readRows(tab, sqler, params)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	sender(w, inter)

}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {

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
