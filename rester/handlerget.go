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
		params, err := service.PrepParam(tab, w, r, fun1)
		if err != nil {
			service.SenderErr(w, err)
			return
		}

		inter, err := service.ReadRows(tab, params)
		if err != nil {
			service.SenderErr(w, err)
			return
		}

		service.Sender(w, inter)
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
