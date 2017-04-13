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

func prepLesen1(tab string, w http.ResponseWriter, r *http.Request, defaults interface{}) (json interface{}, err error) {

	json, err = leser1(w, r, defaults)

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
