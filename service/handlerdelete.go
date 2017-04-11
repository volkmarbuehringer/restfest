package service

import (
	"net/http"
	"restfest/db"
	"strconv"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/gorilla/mux"
)

func deleter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	stmt, err := prepare(tab, db.GenDelete)

	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "delete", err)
		return
	}

	rows := db.DBx.QueryRow(stmt.Name, id)

	inter, err := row1Scanner(tab, rows)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "delete", err)
		return
	}

	sender(w, inter)

}
