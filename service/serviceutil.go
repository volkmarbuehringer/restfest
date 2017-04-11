package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"restfest/db"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
	"github.com/jackc/pgx"
)

var decoder = schema.NewDecoder()

func row1Scanner(tab string, rows *pgx.Row) (stru interface{}, err error) {

	fun := db.ScannerFunMap[tab]

	arr, ts := fun()
	if err = rows.Scan(arr...); err != nil {
		return
	}
	stru = ts
	return

}

func rowScanner(tab string, rows *pgx.Rows) (stru interface{}, err error) {
	t := make([]interface{}, 0)
	fun := db.ScannerFunMap[tab]

	for anz := 0; rows.Next(); anz++ {
		arr, ts := fun()
		if err = rows.Scan(arr...); err != nil {
			return
		}
		t = append(t, ts)
	}
	stru = &t
	return
}

func prepare(tab string, flag db.SQLOper) (stmt *pgx.PreparedStatement, err error) {

	if sqlFun, ok := db.SQLFunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {
		stmt, err = db.DBx.Prepare(tab+strconv.Itoa(int(flag)), sqlFun(flag))

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
		/*
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err = json.NewEncoder(w).Encode(err); err != nil {
				return nil, err
			}
		*/
		return nil, err
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
		/*
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err = json.NewEncoder(w).Encode(err); err != nil {
				return
			}
		*/
		return nil, err
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

func prepLesen(tab string, w http.ResponseWriter, r *http.Request) (json interface{}, err error) {
	if flag, ok := db.FlagMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else if flag == 3 {
		fun1 := db.ParamFunMap[tab]
		json, err = leser1(w, r, fun1())
	} else {
		fun1 := db.EmptyFunMap[tab]
		json, err = leser1(w, r, fun1())
	}

	return
}

func prepParam(tab string, w http.ResponseWriter, r *http.Request) (json interface{}, err error) {
	if fun1, ok := db.ParamFunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {

		err = r.ParseForm()

		if err != nil {
			return
		}

		json = fun1()
		err = decoder.Decode(json, r.Form)

	}

	return
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
