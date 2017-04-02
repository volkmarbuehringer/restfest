package service

import (
	"database/sql"
	"net/http"
	"restfest/db"
	"restfest/gener"

	"github.com/gorilla/mux"
)

var sqlIns = map[string]*sql.Stmt{}

const sqlInsert string = `insert into ` + dbschema + `.%s(%s)values(%s) returning %s`

func poster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]

	json, err := prepLesen(tab, w, r)
	if err != nil {
		senderErr(w, err)
		return
	}

	if _, ok := sqlIns[tab]; !ok {

		if sqlIns[tab], err = prepare(tab, sqlInsert, db.GenInsert); err != nil {
			senderErr(w, err)
			return
		}

	}

	rows, err := sqlIns[tab].Query(gener.ROWInsertFunMap[tab](json)...)
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
