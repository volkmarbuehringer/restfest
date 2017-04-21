package service

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	tab := mux.Vars(r)["tab"]

	params, _, err := prepParam(tab, w, r)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "getall", err)
		return
	}

	inter, err := readRows(tab, params)
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
