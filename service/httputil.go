package service

import (
	"encoding/json"
	"restfest/db"

	"io"
	"io/ioutil"
	"net/http"

	"strconv"
	"strings"

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

func FormReaderS(r *http.Request, name string, defaulter string) string {
	zu := r.FormValue(name)
	zu = strings.Replace(zu, "\"", "", 2)
	if len(zu) > 0 {
		return zu
	}
	return defaulter
}

func Leser1(w http.ResponseWriter, r *http.Request, todo db.PgxGenerIns) error {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err = r.Body.Close(); err != nil {
		return err
	}
	return json.Unmarshal(body, todo)
}

func Sender(w http.ResponseWriter, todos interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

}
