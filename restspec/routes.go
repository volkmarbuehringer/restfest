package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

const pfad string = "/test/service/los/"

var logger = log.New(os.Stdout, "[req] ", 0)

func httpLogger(fn func(w http.ResponseWriter, r *http.Request, p httprouter.Params)) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		logger.Printf("Started %s %s", r.Method, r.URL.Path)
		fn(w, r, p)
		logger.Printf("Completed in %v", time.Since(start))
	}
}

func routing() *httprouter.Router {
	router := httprouter.New()
	router.GET(pfad+":id", httpLogger(getByIDHandlerLos))
	router.GET(pfad, httpLogger(getAllHandlerLos))
	router.POST(pfad, httpLogger(posterLos))
	router.DELETE(pfad+":id", httpLogger(deleterLos))
	router.PUT(pfad+":id", httpLogger(putterLos))

	return router
}
