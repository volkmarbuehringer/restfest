package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restfest/db"
	"restfest/service"

	"github.com/husobee/vestigo"
)

func poster(w http.ResponseWriter, r *http.Request) {

	tab := vestigo.Param(r, "tab")
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

		var data db.PgxGenerIns

		if fun1.Flag == 3 {
			data = fun1.ParamFun()
		} else {
			data = fun1.EmptyInsFun()
		}
		err = json.NewDecoder(io.LimitReader(r.Body, 1048576)).Decode(data)

		if err != nil {
			service.SenderErr(w, err)
			return
		}
		input := data.ROWInsert()

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
