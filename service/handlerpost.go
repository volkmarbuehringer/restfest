package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"restfest/gener"

	"github.com/gorilla/mux"
)

var sqlIns = map[string]*sql.Stmt{}

const sqlInsert string = `insert into ` + dbschema + `.%s(%s)values(%s) returning %s`

func poster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]

	if fun1, ok := gener.EmptyFunMap[tab]; !ok {
		senderErr(w, fmt.Errorf("Tabelle nicht gefunden: %s", tab))
		return
	} else {
		json, err := leser1(w, r, fun1()[0])
		if err != nil {
			senderErr(w, err)
			return
		}

		if _, ok := sqlIns[tab]; !ok {

			if sqlIns[tab], err = prepare(tab, sqlInsert, gener.GenInsert); err != nil {
				senderErr(w, err)
				return
			}

		}
		x := gener.ROWInsertFunMap[tab](json)
		rows, err := sqlIns[tab].Query(x...)
		if err != nil {
			senderErr(w, err)
			return
		}

		inter, err := gener.ROWSFunMap[tab](rows, 1)
		if err != nil {
			senderErr(w, err)
			return
		}

		sender(w, inter[0])
	}
}
