package service

import (
	"net/http"
	"restfest/db"
	"strconv"

	"github.com/gorilla/mux"
)

func deleter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	stmt, err := prepare(tab, db.GenDelete)

	if err != nil {
		senderErr(w, err)
		return
	}
	empty := []interface{}{id}

	rows, err := db.DBx.Query(stmt.Name, empty...)
	if err != nil {
		senderErr(w, err)
		return
	}

	inter, err := rowScanner(tab, rows, 1)
	if err != nil {
		senderErr(w, err)
		return
	}

	sender(w, inter)

}
