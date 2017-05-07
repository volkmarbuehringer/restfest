package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/husobee/vestigo"
)

const pfad string = "/test/service/los"

var logger = log.New(os.Stdout, "[req] ", 0)

func httpLogger(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		logger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// Index shows the profile index.
func Index(w http.ResponseWriter, r *http.Request) {
	pprof.Index(w, r)
}

// Profile shows the individual profiles.
func Profile(w http.ResponseWriter, r *http.Request) {
	switch vestigo.Param(r, "pprof") {
	case "cmdline":
		pprof.Cmdline(w, r)
	case "profile":
		pprof.Profile(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	case "trace":
		pprof.Trace(w, r)
	default:
		Index(w, r)
	}
}

func routing() *vestigo.Router {
	// you can enable trace by setting this to true
	vestigo.AllowTrace = true

	router := vestigo.NewRouter()
	router.Get(pfad+"/:id", httpLogger(getByIDHandlerLos))
	router.Get(pfad, httpLogger(getAllHandlerLos))
	router.Post(pfad, httpLogger(posterLos))
	router.Delete(pfad+"/:id", httpLogger(deleterLos))
	router.Put(pfad+"/:id", httpLogger(putterLos))

	router.Get("/debug/pprof/", Index)
	router.Get("/debug/pprof/:pprof", Profile)

	return router
}
