package service

import (
	"encoding/json"
	"restfest/db"

	"net/http"

	log15 "gopkg.in/inconshreveable/log15.v2"
)

func PrepParam(params db.PgxGenerIns, w http.ResponseWriter, r *http.Request) error {

	err := r.ParseForm()
	if err != nil {
		return err
	}

	inter := params.Scanner()

	return inter.ConvertAtoI(params.Reader(r.Form))

}

func SenderErr(w http.ResponseWriter, err error) {
	type JSONErr struct {
		Error string
	}
	log15.Error("DBFehler", "getall", err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err1 := json.NewEncoder(w).Encode(JSONErr{err.Error()}); err1 != nil {
		panic(err1)
	}

}

/*
func FormReader(r *http.Request, name string, defaulter int) int {
	zu := r.FormValue(name)
	if len(zu) > 0 {
		la, err := strconv.Atoi(zu)
		if err == nil {
			return la
		}
	}
	return defaulter
}
*/

func Sender(w http.ResponseWriter, todos interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

}
