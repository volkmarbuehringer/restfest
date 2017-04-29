package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/service"
	"strconv"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/gorilla/mux"
)

func putter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	tab := mux.Vars(r)["tab"]

	if sqlFun, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		service.SenderErr(w, err)
		log15.Error("DBFehler", " put", err)
		return
	} else {
		tx, err := db.DBx.Begin()
		if err != nil {
			service.SenderErr(w, err)
			log15.Error("DBFehler", "begin put", err)
			return
		}
		defer tx.Rollback()
		params := sqlFun.ParamFun()
		sql := params.SQL(db.GenSelectID)
		rows := tx.QueryRow(sql+" for update", id)

		inter := sqlFun.EmptyFun()

		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, err)
			log15.Error("DBFehler", "prep put", err)
			return
		}

		err = service.Leser1(w, r, inter)
		if err != nil {
			service.SenderErr(w, err)
			log15.Error("DBFehler", "put", err)
			return
		}

		x := append(inter.ROWInsert(), id)

		fmt.Println(x)
		sql = inter.SQL(db.GenUpdate)
		rows = tx.QueryRow(sql, x...)
		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, err)
			log15.Error("DBFehler", "update", err)
			return
		}

		err = tx.Commit()
		if err != nil {
			service.SenderErr(w, err)
			log15.Error("DBFehler", "commit", err)
			return
		}

		service.Sender(w, inter)

	}
}
