package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

const pfad string = "/test/service/:tab/"

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

	router.GET(pfad+":id", httpLogger(getByIDHandler))
	router.GET(pfad, httpLogger(getAllHandler))
	router.POST(pfad, httpLogger(poster))
	router.DELETE(pfad+":id", httpLogger(deleter))
	router.PUT(pfad+":id", httpLogger(putter))

	return router
}
