package service

import (
	"net/http"
	"restfest/db"
	"strconv"

	log15 "gopkg.in/inconshreveable/log15.v2"

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
		log15.Error("DBFehler", "put", err)
		return
	}

	var stmt *pgx.PreparedStatement
	if stmt, err = prepare(tab, db.GenUpdate); err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "put", err)
		return
	}

	x := db.ROWInsertFunMap[tab](json)
	x = append(x, id)
	rows := db.DBx.QueryRow(stmt.Name, x...)

	inter, err := row1Scanner(tab, rows)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "put", err)
		return
	}

	sender(w, inter)

}
