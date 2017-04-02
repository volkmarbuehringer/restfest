package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"restfest/db"
	"restfest/gener"
)

const dbschema = "public"

var sqlAll = map[string]*sql.Stmt{}
var sqlID = map[string]*sql.Stmt{}

func getByIDHandler3(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]
	ga := db.Pager{}
	ga.Length = formReader(r, "length", 100)
	ga.Offset = formReader(r, "offset", 0)
	ga.Where = formReaderS(r, "where", "")

	if _, ok := sqlAll[tab]; !ok {
		if sqlFun, ok := gener.SQLFunMap[tab]; !ok {
			senderErr(w, fmt.Errorf("Tabelle nicht gefunden: %s", tab))
			return
		} else {
			sql, sql1 := sqlFun()
			sqls := fmt.Sprintf("select %s from "+dbschema+".%s order by %s limit $1 offset $2", strings.Join(sql, ","), tab, sql1)
			var err error
			fmt.Println("prep", sqls)
			if sqlAll[tab], err = db.DB.Prepare(sqls); err != nil {
				senderErr(w, err)
				return
			}

		}
	}
	rows, err := sqlAll[tab].Query(ga.Length, ga.Offset)
	if err != nil {
		senderErr(w, err)
		return
	}
	defer rows.Close()

	inter, err := gener.ROWSFunMap[tab](rows, 0)
	if err != nil {
		senderErr(w, err)
		return
	}

	sender(w, inter[0])

}

func getByIDHandler5(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]
	id, _ := strconv.Atoi(vars["id"])

	if _, ok := sqlID[tab]; !ok {
		if sqlFun, ok := gener.SQLFunMap[tab]; !ok {
			senderErr(w, fmt.Errorf("Tabelle nicht gefunden: %s", tab))
			return
		} else {
			sql, sql1 := sqlFun()
			sqls := fmt.Sprintf("select %s from "+dbschema+".%s where %s=$1", strings.Join(sql, ","), tab, sql1)

			fmt.Println("prep", sqls)
			var err error
			if sqlID[tab], err = db.DB.Prepare(sqls); err != nil {
				senderErr(w, err)
				return
			}
		}
	}
	rows, err := sqlID[tab].Query(id)
	if err != nil {
		senderErr(w, err)
		return
	}
	defer rows.Close()

	if inter, err := gener.ROWSFunMap[tab](rows, 1); err != nil {
		senderErr(w, err)
		return
	} else {
		sender(w, inter[0])
	}

}
