package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/gener"
	"strconv"
	"strings"

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

			if sqlFun, ok := gener.SQLUpdateFunMap[tab]; !ok {
				senderErr(w, fmt.Errorf("Tabelle nicht gefunden: %s", tab))
				return
			} else {

				sql1, sql2, sql3 := sqlFun()
				sqls := fmt.Sprintf(sqlUpdate, tab, strings.Join(sql1, ","), sql2,
					strings.Join(sql3, ","))

				//rows, err := m.RowInsert(&json, tab)
				fmt.Println(tab, sqls)

				if sqlUpd[tab], err = db.DB.Prepare(sqls); err != nil {
					senderErr(w, err)
					return
				}
			}
		}

		rows, err := gener.ROWUpdateFunMap[tab](sqlUpd[tab], json, id)

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
