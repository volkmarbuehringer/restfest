package main

import (
	"net/http"
	"restfest/db"
	"restfest/restspec/gener"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
)

func deleterLos(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	var los = gener.Los{}
	stmt, err := service.PrepareSQL("losDEL", func() string {
		return "delete from los where id =$1 returning " + gener.LosSQL.All
	})

	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	rows := db.DBx.QueryRow(stmt.Name, id)

	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	service.Sender(w, los)

}
