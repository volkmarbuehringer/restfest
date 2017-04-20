package service

import (
	"fmt"
	"net/http"
	"restfest/db"
	"strconv"

	"github.com/gorilla/mux"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]

	params, funMap, err := prepParam(tab, w, r)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	var sqler db.SQLOper
	switch funMap.Flag {
	case 3:
		sqler = db.GenFunction
	case 4:
		sqler = db.GenSelectAll1
	case 1, 2:
		sqler = db.GenSelectAll
	default:
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		senderErr(w, err)
		return

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
	id, _ := strconv.Atoi(vars["id"])

	if inter, err := readRow(vars["tab"], id); err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	} else {
		sender(w, inter)
	}

}
