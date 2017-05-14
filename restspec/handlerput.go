package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restfest/db"
	"restfest/restspec/gener"
	"restfest/service"
	"strconv"

	"github.com/husobee/vestigo"
)

func putterLos(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(vestigo.Param(r, "id"))
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	tx, err := db.DBx.Begin()
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	defer tx.Rollback()
	var los gener.Los

	rows := tx.QueryRow(fmt.Sprintf("select %s from "+db.DBschema+".los where id = $1 for update",
		gener.LosSQL.All,
	), id)

	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, r, err)
		return
	}
	err = json.NewDecoder(io.LimitReader(r.Body, 1048)).Decode(&los)

	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	x := append(los.ROWInsert(), id)

	sql := fmt.Sprintf(`update `+db.DBschema+`.los set %s where %s returning %s`,
		gener.LosSQL.BindsUpdate, gener.LosSQL.PKUpdate, gener.LosSQL.All)
	rows = tx.QueryRow(sql, x...)
	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		service.SenderErr(w, r, err)
		return
	}

	service.Sender(w, los)

}
