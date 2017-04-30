package main

import (
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
)

func deleterWeburl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var weburl = generrestspec.Weburl{}
	stmt, err := service.Prepare("weburl", db.GenDelete, &weburl)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	rows := db.DBx.QueryRow(stmt.Name, id)

	err = rows.Scan(weburl.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	service.Sender(w, weburl)

}
