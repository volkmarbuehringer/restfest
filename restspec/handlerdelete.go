package main

import (
	"net/http"
	"restfest/db"
	gener "restfest/restspec/gener"
	"restfest/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func deleterLos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	var los = gener.Los{}
	stmt, err := service.PrepareSQL("losDEL", func() string {
		return "delete from los where id =$1 returning " + gener.LosSQL.All
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
