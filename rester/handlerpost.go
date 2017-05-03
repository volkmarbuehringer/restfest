package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/service"

	"github.com/julienschmidt/httprouter"
)

func poster(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	tab := ps.ByName("tab")
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			service.SenderErr(w, err)
			return
		}
	} else {
		stmt, err := service.Prepare(tab, service.GetSqlStmt(-1, fun1.Flag), fun1.EmptyFun())
		if err != nil {
			service.SenderErr(w, err)
			return
		}

		var json db.PgxGenerIns

		if fun1.Flag == 3 {
			json = fun1.ParamFun()
		} else {
			json = fun1.EmptyInsFun()
		}
		err = service.Leser1(w, r, json)

		if err != nil {
			service.SenderErr(w, err)
			return
		}
		input := json.ROWInsert()

		rows := db.DBx.QueryRow(stmt.Name, input...)

		inter := fun1.EmptyFun()

		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, err)
			return
		}

		service.Sender(w, inter)
	}
}
