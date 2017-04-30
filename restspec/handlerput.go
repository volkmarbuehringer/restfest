package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
	"strconv"

	"github.com/gorilla/mux"
)

func putterWeburl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tx, err := db.DBx.Begin()
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	defer tx.Rollback()
	var params generrestspec.WeburlParams
	var weburl generrestspec.Weburl
	sql := params.SQL(db.GenSelectID)
	rows := tx.QueryRow(sql+" for update", id)

	err = rows.Scan(weburl.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	err = service.Leser1(w, r, &weburl)
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	x := append(weburl.ROWInsert(), id)

	fmt.Println(x)
	sql = weburl.SQL(db.GenUpdate)
	rows = tx.QueryRow(sql, x...)
	err = rows.Scan(weburl.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	service.Sender(w, weburl)

}
