package service

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"restfest/db"
	"restfest/gener"
)

const dbschema = "public"

var sqlAll = map[string]*sql.Stmt{}
var sqlID = map[string]*sql.Stmt{}

const selectAll = "select %s from " + dbschema + ".%s order by %s limit $1 offset $2"
const selectID = "select %s from " + dbschema + ".%s where %s=$1"

func getByIDHandler3(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tab := vars["tab"]
	ga := db.Pager{}
	ga.Length = formReader(r, "length", 100)
	ga.Offset = formReader(r, "offset", 0)
	ga.Where = formReaderS(r, "where", "")

	if _, ok := sqlAll[tab]; !ok {
		var err error
		if sqlAll[tab], err = prepare(tab, selectAll, db.GenSelect); err != nil {
			senderErr(w, err)
			return
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
		var err error
		if sqlID[tab], err = prepare(tab, selectID, db.GenSelect); err != nil {
			senderErr(w, err)
			return
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
