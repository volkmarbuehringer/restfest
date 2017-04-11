package service

import (
	"net/http"
	"restfest/db"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

func poster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]

	json, err := prepLesen(tab, w, r)
	if err != nil {
		senderErr(w, err)
		return
	}
	var stmt *pgx.PreparedStatement

	if db.FlagMap[tab] == 3 {
		stmt, err = prepare(tab, db.GenFunction)
	} else {
		stmt, err = prepare(tab, db.GenInsert)
	}
	if err != nil {
		senderErr(w, err)
		return
	}

	var input []interface{}
	if db.FlagMap[tab] == 3 {
		input = db.ROWQueryFunMap[tab](json)
	} else {
		input = db.ROWInsertFunMap[tab](json)
	}
	rows := db.DBx.QueryRow(stmt.Name, input...)

	inter, err := row1Scanner(tab, rows)
	if err != nil {
		senderErr(w, err)
		return
	}

	sender(w, inter)

}
