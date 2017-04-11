package service

import (
	"net/http"
	"restfest/db"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

func putter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	json, err := prepLesen(tab, w, r)
	if err != nil {
		senderErr(w, err)
		return
	}

	var stmt *pgx.PreparedStatement
	if stmt, err = prepare(tab, db.GenUpdate); err != nil {
		senderErr(w, err)
		return
	}

	x := db.ROWInsertFunMap[tab](json)
	x = append(x, id)
	rows, err := db.DBx.Query(stmt.Name, x...)

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
