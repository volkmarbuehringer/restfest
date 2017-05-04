package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	gener "restfest/restspec/gener"
	"restfest/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func putterLos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	tx, err := db.DBx.Begin()
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	defer tx.Rollback()
	var los gener.Los

	rows := tx.QueryRow(fmt.Sprintf("select %s from "+db.DBschema+".los where id = $1 for update",
		gener.LosSQL.All,
	), id)

	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	err = service.Leser1(w, r, &los)
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	x := append(los.ROWInsert(), id)

	fmt.Println(x)

	sql := fmt.Sprintf(`update `+db.DBschema+`.los set %s where %s returning %s`,
		gener.LosSQL.BindsUpdate, gener.LosSQL.PKUpdate, gener.LosSQL.All)
	rows = tx.QueryRow(sql, x...)
	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	service.Sender(w, los)

}
