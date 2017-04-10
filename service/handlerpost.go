package service

import (
	"database/sql"
	"net/http"
	"restfest/db"

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

		if db.FlagMap[tab] == 3 {
			sqlIns[tab], err = prepare(tab, selectFun, db.GenSelect)
		} else {
			sqlIns[tab], err = prepare(tab, sqlInsert, db.GenInsert)
		}
		if err != nil {
			senderErr(w, err)
			return
		}

	}
	var input []interface{}
	if db.FlagMap[tab] == 3 {
		input = db.ROWQueryFunMap[tab](json)
	} else {
		input = db.ROWInsertFunMap[tab](json)
	}

	rows, err := sqlIns[tab].Query(input...)
	if err != nil {
		senderErr(w, err)
		return
	}

	inter, err := rowScanner(tab, rows, 1)
	if err != nil {
		senderErr(w, err)
		return
	}

	sender(w, inter)

}
