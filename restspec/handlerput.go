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

func putterLos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tx, err := db.DBx.Begin()
	if err != nil {
		service.SenderErr(w, err)
		return
	}
	defer tx.Rollback()
	var los generrestspec.Los

	rows := tx.QueryRow(fmt.Sprintf("select %s from "+db.DBschema+".los where id = $1 for update",
		generrestspec.LosSQL.All,
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
		generrestspec.LosSQL.BindsUpdate, generrestspec.LosSQL.PKUpdate, generrestspec.LosSQL.All)
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
