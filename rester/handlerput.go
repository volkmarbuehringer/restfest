package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
)

func putter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	tab := vestigo.Param(r, "tab")
	if sqlFun, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		service.SenderErr(w, err)
		return
	} else {
		tx, err := db.DBx.Begin()
		if err != nil {
			service.SenderErr(w, err)
			return
		}
		defer tx.Rollback()
		params := sqlFun.ParamFun()
		sql := params.SQL(db.GenSelectID)
		rows := tx.QueryRow(sql+" for update", id)

		inter := sqlFun.EmptyFun()

		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, err)
			return
		}

		err = service.Leser1(w, r, inter)
		if err != nil {
			service.SenderErr(w, err)
			return
		}

		x := append(inter.ROWInsert(), id)

		fmt.Println(x)
		sql = inter.SQL(db.GenUpdate)
		rows = tx.QueryRow(sql, x...)
		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, err)
			return
		}

		err = tx.Commit()
		if err != nil {
			service.SenderErr(w, err)
			return
		}

		service.Sender(w, inter)

	}
}
