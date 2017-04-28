package service

import (
	"fmt"
	"net/http"
	"restfest/db"
	"restfest/generrester"
	"strconv"

	"github.com/gorilla/mux"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

func getAllHandler(w http.ResponseWriter, r *http.Request) {

	tab := mux.Vars(r)["tab"]
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}

	} else {
		params, err := prepParam(tab, w, r, fun1)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}

		inter, err := readRows(tab, params)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}

		sender(w, inter)
	}
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	tab := vars["tab"]
	if fun1, ok := db.FunMap[tab]; !ok {
		err := fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		}
	} else {
		if inter, err := readRow(vars["tab"], id, fun1.ParamFun()); err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "getall", err)
			return
		} else {
			sender(w, inter)
		}
	}
}

func getByIDHandlerWeburl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	params := new(generrester.WeburlParams)
	if stmt, _, err := prepare("weburl", db.GenSelectID, params); err != nil {
		senderErr(w, err)
		log15.Error("DBFehler", "weburl", err)
		return
	} else {

		row := db.DBx.QueryRow(stmt.Name, id)

		weburl := new(generrester.Weburl)

		err = row.Scan(weburl.Scanner()...)
		if err != nil {
			senderErr(w, err)
			log15.Error("DBFehler", "weburl", err)
			return
		} else {
			*weburl.Zusatz = 333 //set value in struct
			sender(w, weburl)

		}

	}

}
