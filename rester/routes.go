package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/husobee/vestigo"
)

const pfad string = "/test/service/:tab"

var logger = log.New(os.Stdout, "[req] ", 0)

func httpLogger(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		logger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}
func routing() *vestigo.Router {
	router := vestigo.NewRouter()

	router.Get(pfad, httpLogger(getAllHandler))
	router.Get(pfad+"/:id", httpLogger(getByIDHandler))
	router.Post(pfad, httpLogger(poster))
	router.Delete(pfad+"/:id", httpLogger(deleter))
	router.Put(pfad+"/:id", httpLogger(putter))

	return router
}
