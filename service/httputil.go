package service

import (
	"encoding/json"
	"restfest/db"

	"net/http"
)

func PrepParam(params db.PgxGenerIns, w http.ResponseWriter, r *http.Request) error {

	err := r.ParseForm()
	if err != nil {
		return err
	}

	inter := params.Scanner()

	return inter.ConvertAtoI(params.Reader(r.Form))

}

func SenderErr(w http.ResponseWriter, r *http.Request, err error) {
	type JSONErr struct {
		Error string
	}
	l := MustFromContext(r.Context())

	l.Error(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err1 := json.NewEncoder(w).Encode(JSONErr{err.Error()}); err1 != nil {
		panic(err1)
	}

}

func Sender(w http.ResponseWriter, todos interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

}
