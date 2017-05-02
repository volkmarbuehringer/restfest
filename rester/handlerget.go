package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	tab := mux.Vars(r)["tab"]
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			service.SenderErr(w, err)
			return
		}

	} else {
		params := fun1.ParamFun()
		err := service.PrepParam(params, w, r)
		if err != nil {
			service.SenderErr(w, err)
			return
		}

		rows, err := service.ReadRows(tab, params)
		if err != nil {
			service.SenderErr(w, err)
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
				service.SenderErr(w, err)
				return
			}

		}
		if err = iter.Err(); err != nil {
			service.SenderErr(w, err)
			return
		}

		w.Write([]byte("]"))

	}
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	tab := vars["tab"]
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			service.SenderErr(w, err)
			return
		}
	} else {
		if inter, err := service.ReadRow(vars["tab"], id, fun1.ParamFun()); err != nil {
			service.SenderErr(w, err)
			return
		} else {
			service.Sender(w, inter)
		}
	}
}
