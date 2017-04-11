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

	rows := db.DBx.QueryRow(stmt.Name, id)

	inter, err := row1Scanner(tab, rows)
	if err != nil {
		senderErr(w, err)
		return
	}

	sender(w, inter)

}
