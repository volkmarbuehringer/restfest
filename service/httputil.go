package service

import (
	"encoding/json"
	"fmt"
	"restfest/db"

	"io"
	"io/ioutil"
	"net/http"

	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func prepParam(tab string, w http.ResponseWriter, r *http.Request) (json db.PgxGenerIns, fun1 db.TFunMap, sqler db.SQLOper, err error) {
	var ok bool
	if fun1, ok = db.FunMap[tab]; !ok {
		err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
		return
	} else {

		switch fun1.Flag {
		case 3:
			sqler = db.GenFunction
		case 4:
			sqler = db.GenSelectAll1
		case 1, 2:
			sqler = db.GenSelectAll
		default:
			err = fmt.Errorf("Tabelle nicht gefunden: %s", tab)
			return

		}
		err = r.ParseForm()

		if err != nil {
			return
		}

		json = fun1.ParamFun()
		err = decoder.Decode(json, r.Form)

	}

	return
}

func prepLesen(flag db.TFunMap, w http.ResponseWriter, r *http.Request) (json db.PgxGenerIns, err error) {
	if flag.Flag == 3 {
		json, err = leser1(w, r, flag.ParamFun())
	} else {
		json, err = leser1(w, r, flag.EmptyInsFun())
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

func leser1(w http.ResponseWriter, r *http.Request, todo db.PgxGenerIns) (db.PgxGenerIns, error) {

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
