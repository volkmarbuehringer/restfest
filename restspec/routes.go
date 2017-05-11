package main

import (
	"net/http"
	"net/http/pprof"

	"github.com/husobee/vestigo"
)

const pfad string = "/test/service/los"

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

	router.Get(pfad+"/:id", getByIDHandlerLos)
	router.Get(pfad, getAllHandlerLos)
	router.Post(pfad, posterLos)
	router.Delete(pfad+"/:id", deleterLos)
	router.Put(pfad+"/:id", putterLos)

	router.Get("/debug/pprof/", Index)
	router.Get("/debug/pprof/:pprof", Profile)

	return router
}
