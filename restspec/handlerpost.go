package main

import (
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"
)

func posterWeburl(w http.ResponseWriter, r *http.Request) {

	var weburl = generrestspec.Weburl{}

	stmt, err := service.Prepare("weburl", service.GetSqlStmt(-1, 1), &weburl)
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	err = service.Leser1(w, r, &weburl)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	rows := db.DBx.QueryRow(stmt.Name, weburl.ROWInsert()...)

	err = rows.Scan(weburl.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	service.Sender(w, weburl)

}
