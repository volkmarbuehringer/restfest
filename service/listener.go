package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Listen(routes *httprouter.Router) {

	srv := &http.Server{
		ReadTimeout:  200 * time.Second,
		WriteTimeout: 20000 * time.Second,
		Addr:         ":8080",
		Handler:      routes,
	}
	fmt.Println("listen")
	log.Fatal(srv.ListenAndServe())
}
