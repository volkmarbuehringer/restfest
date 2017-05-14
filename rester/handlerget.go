package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	tab := vestigo.Param(r, "tab")
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

	} else {
		params := fun1.ParamFun()
		err := service.PrepParam(params, w, r)
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

		rows, err := ReadRows(tab, params)
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}

		defer rows.Close()
		iter := fun1.Iterator()
		iter.NewCopy(rows)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte("["))
		var trenner string
		for anz := 0; iter.Next(); anz++ {

			if anz > 0 {
				trenner = ","
			}
			err = iter.Value().Writer(w, trenner)
			if err != nil {
				service.SenderErr(w, r, err)
				return
			}

		}
		if err = iter.Err(); err != nil {
			service.SenderErr(w, r, err)
			return
		}

		w.Write([]byte("]"))

	}
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	tab := vestigo.Param(r, "tab")
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			service.SenderErr(w, r, err)
			return
		}
	} else {
		if inter, err := ReadRow(tab, id, fun1.ParamFun()); err != nil {
			service.SenderErr(w, r, err)
			return
		} else {
			service.Sender(w, inter)
		}
	}
}
