package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
)

func deleter(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	tab := vestigo.Param(r, "tab")
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			service.SenderErr(w, err)
			return
		}
	} else {
		stmt, err := service.Prepare(tab, db.GenDelete, fun1.EmptyFun())

		if err != nil {
			service.SenderErr(w, err)
			return
		}

		rows := db.DBx.QueryRow(stmt.Name, id)

		inter := fun1.EmptyFun()

		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, err)
			return
		}

		service.Sender(w, inter)
	}
}
