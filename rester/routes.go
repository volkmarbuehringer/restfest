package main

import "github.com/husobee/vestigo"

const pfad string = "/test/service/:tab"

func routing() *vestigo.Router {
	router := vestigo.NewRouter()

	router.Get(pfad, getAllHandler)
	router.Get(pfad+"/:id", getByIDHandler)
	router.Post(pfad, poster)
	router.Delete(pfad+"/:id", deleter)
	router.Put(pfad+"/:id", putter)

	return router
}
