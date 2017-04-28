package service

import (
	"fmt"
	"net/http"
	"restfest/db"

	"github.com/gorilla/mux"
	"gopkg.in/inconshreveable/log15.v2"
)

func poster(w http.ResponseWriter, r *http.Request) {

	tab := mux.Vars(r)["tab"]
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}
	} else {
		stmt, tmap, err := prepare(tab, -1, fun1.EmptyFun())
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "post", err)
			return
		}

		var json db.PgxGenerIns

		if tmap.Flag == 3 {
			json = tmap.ParamFun()
		} else {
			json = tmap.EmptyInsFun()
		}
		err = leser1(w, r, json)

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
}
