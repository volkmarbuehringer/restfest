package service

import (
	"fmt"
	"net/http"
	"restfest/db"
	"strconv"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/gorilla/mux"
)

func putter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	if sqlFun, ok := db.SQLFunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		senderErr(w, err)
		log15.Error("DBFehler", " put", err)
		return
	} else {
		tx, err := db.DBx.Begin()
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "begin put", err)
			return
		}
		defer tx.Rollback()
		rows := tx.QueryRow(sqlFun(db.GenSelectID)+" for update", id)
		var inter interface{}
		inter, err = row1Scanner(tab, rows)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "prep put", err)
			return
		}

		json, err := prepLesen1(tab, w, r, inter)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "put", err)
			return
		}

		x := db.ROWInsertFunMap[tab](json)
		x = append(x, id)

		fmt.Println(x)
		rows = tx.QueryRow(sqlFun(db.GenUpdate), x...)

		inter, err = row1Scanner(tab, rows)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "update", err)
			return
		}

		err = tx.Commit()
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "commit", err)
			return
		}

		sender(w, inter)

	}
}
