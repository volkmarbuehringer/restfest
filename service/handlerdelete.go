package service

import (
	"fmt"
	"net/http"
	"restfest/db"
	"strconv"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/gorilla/mux"
)

func deleter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	tab := mux.Vars(r)["tab"]
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}
	} else {
		stmt, funMap, err := prepare(tab, db.GenDelete, fun1.EmptyFun())

		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "delete", err)
			return
		}

		rows := db.DBx.QueryRow(stmt.Name, id)

		inter := funMap.EmptyFun()

		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "delete", err)
			return
		}

		sender(w, inter)
	}
}
