package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"restfest/gener"
	"strconv"
	"strings"

	"restfest/db"
)

func prepare(tab string, sqlSt string, flag gener.SQLOper) (stmt *sql.Stmt, err error) {
	if sqlFun, ok := gener.SQLFunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {
		sqls := fmt.Sprintf(sqlSt, sqlFun(tab, flag)...)
		fmt.Println("prep", sqls)
		stmt, err = db.DB.Prepare(sqls)

	}
	return
}

func formReader(r *http.Request, name string, defaulter int) int {
	zu := r.FormValue(name)
	if len(zu) > 0 {
		la, err := strconv.Atoi(zu)
		if err == nil {
			return la
		}
	}
	return defaulter
}

func formReaderS(r *http.Request, name string, defaulter string) string {
	zu := r.FormValue(name)
	zu = strings.Replace(zu, "\"", "", 2)
	if len(zu) > 0 {
		return zu
	}
	return defaulter
}

func leser1(w http.ResponseWriter, r *http.Request, todo interface{}) (interface{}, error) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, err
	}
	if err = r.Body.Close(); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return nil, err
		}
	}
	return todo, err
}

func leser(w http.ResponseWriter, r *http.Request) (todo map[string]interface{}, err error) {

	todo = make(map[string]interface{})

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return
	}
	if err = r.Body.Close(); err != nil {
		return
	}
	if err = json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
	}
	return
}

func sender(w http.ResponseWriter, todos interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

}

func senderErr(w http.ResponseWriter, err error) {
	type JSONErr struct {
		Error string
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err1 := json.NewEncoder(w).Encode(JSONErr{err.Error()}); err1 != nil {
		panic(err1)
	}

}
