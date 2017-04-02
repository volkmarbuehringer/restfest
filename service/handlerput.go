package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"restfest/gener"
	"strconv"

	"github.com/gorilla/mux"
)

var sqlUpd = map[string]*sql.Stmt{}

const sqlUpdate string = `update ` + dbschema + `.%s set %s where %s returning %s`

func putter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	if fun1, ok := gener.EmptyFunMap[tab]; !ok {
		senderErr(w, fmt.Errorf("Tabelle nicht gefunden: %s", tab))
		return
	} else {

		json, err := leser1(w, r, fun1()[0])
		if err != nil {
			senderErr(w, err)
			return
		}

		if _, ok := sqlUpd[tab]; !ok {

			if sqlUpd[tab], err = prepare(tab, sqlUpdate, gener.GenUpdate); err != nil {
				senderErr(w, err)
				return
			}

		}
		x := gener.ROWInsertFunMap[tab](json)
		x = append(x, id)
		rows, err := sqlUpd[tab].Query(x...)

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
