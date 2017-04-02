package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/gener"
	"strings"

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

			if sqlFun, ok := gener.SQLInsertFunMap[tab]; !ok {
				senderErr(w, fmt.Errorf("Tabelle nicht gefunden: %s", tab))
			} else {

				sql1, sql2, sql3 := sqlFun()
				sqls := fmt.Sprintf(sqlInsert, tab, strings.Join(sql1, ","), strings.Join(sql2, ","),
					strings.Join(sql3, ","))

				//rows, err := m.RowInsert(&json, tab)
				fmt.Println(tab, sqls)

				if sqlIns[tab], err = db.DB.Prepare(sqls); err != nil {
					senderErr(w, err)
					return
				}
			}
		}
		rows, err := gener.ROWInsertFunMap[tab](sqlIns[tab], json)

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
