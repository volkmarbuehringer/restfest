package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/generrestspec"
	"restfest/service"

	"github.com/julienschmidt/httprouter"
)

func posterLos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var los = generrestspec.Los{L_iban: db.String("default")}

	stmt, err := service.PrepareSQL("losInsert", func() string {
		return fmt.Sprintf(`insert into `+db.DBschema+`.los(%s)values(%s) returning %s`,
			generrestspec.LosSQL.Inserts, generrestspec.LosSQL.BindsInsert, generrestspec.LosSQL.All)
	})
	if err != nil {
		service.SenderErr(w, err)
		return
	}

	err = service.Leser1(w, r, &los)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	rows := db.DBx.QueryRow(stmt.Name, los.ROWInsert()...)

	err = rows.Scan(los.Scanner()...)

	if err != nil {
		service.SenderErr(w, err)
		return
	}

	service.Sender(w, los)

}
