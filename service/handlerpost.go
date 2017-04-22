package service

import (
	"net/http"
	"restfest/db"

	"github.com/gorilla/mux"
	"gopkg.in/inconshreveable/log15.v2"
)

func poster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

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

	inter := tmap.EmptyFun()

	err = rows.Scan(inter.Scanner()...)

	if err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "post", err)
		return
	}

	sender(w, inter)

}
