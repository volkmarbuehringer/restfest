package service

import (
	"net/http"
	"restfest/db"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
)

func poster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var stmt *pgx.PreparedStatement
	stmt, tmap, err := prepare(vars["tab"], -1)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "post", err)
		return
	}

	json, err := prepLesen(tmap, w, r)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "post", err)
		return
	}
	input := json.ROWInsert()

	rows := db.DBx.QueryRow(stmt.Name, input...)

	inter, err := row1Scanner(tmap, rows)
	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "post", err)
		return
	}

	sender(w, inter)

}
