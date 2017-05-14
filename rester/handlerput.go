package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restfest/db"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
)

func putter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	tab := vestigo.Param(r, "tab")
	if sqlFun, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		service.SenderErr(w, r, err)
		return
	} else {
		tx, err := db.DBx.Begin()
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}
		defer tx.Rollback()
		params := sqlFun.ParamFun()
		sql := params.SQL(db.GenSelectID)
		rows := tx.QueryRow(sql+" for update", id)

		inter := sqlFun.EmptyFun()

		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

		err = json.NewDecoder(io.LimitReader(r.Body, 1048576)).Decode(inter)
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

		x := append(inter.ROWInsert(), id)

		fmt.Println(x)
		sql = inter.SQL(db.GenUpdate)
		rows = tx.QueryRow(sql, x...)
		err = rows.Scan(inter.Scanner()...)

		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

		err = tx.Commit()
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

		service.Sender(w, inter)

	}
}
