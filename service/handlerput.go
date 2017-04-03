package service

import (
	"database/sql"
	"net/http"
	"restfest/db"
	"strconv"

	"github.com/gorilla/mux"
)

var sqlUpd = map[string]*sql.Stmt{}

const sqlUpdate string = `update ` + dbschema + `.%s set %s where %s returning %s`

func putter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	json, err := prepLesen(tab, w, r)
	if err != nil {
		senderErr(w, err)
		return
	}

	if _, ok := sqlUpd[tab]; !ok {

		if sqlUpd[tab], err = prepare(tab, sqlUpdate, db.GenUpdate); err != nil {
			senderErr(w, err)
			return
		}

	}

	x := db.ROWInsertFunMap[tab](json)
	x = append(x, id)
	rows, err := sqlUpd[tab].Query(x...)

	if err != nil {
		senderErr(w, err)
		return
	}

	inter, err := rowScanner(tab, rows, 1)
	if err != nil {
		senderErr(w, err)
		return
	}
	sender(w, inter[0])

}
