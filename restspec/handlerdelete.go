package main

import (
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
)

func deleterLos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var los = generrestspec.Los{}
	stmt, err := service.PrepareSQL("losDEL", func() string {
		return "delete from los where id =$1 returning " + generrestspec.LosSQL.All
	})

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	rows := db.DBx.QueryRow(stmt.Name, id)

	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	service.Sender(w, los)

}
