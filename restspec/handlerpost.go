package main

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/restspec/gener"
	"restfest/service"
)

func posterLos(w http.ResponseWriter, r *http.Request) {

	var los = gener.Los{L_iban: db.String("default")}

	stmt, err := service.PrepareSQL("losInsert", func() string {
		return fmt.Sprintf(`insert into `+db.DBschema+`.los(%s)values(%s) returning %s`,
			gener.LosSQL.Inserts, gener.LosSQL.BindsInsert, gener.LosSQL.All)
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
