package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func Listen(routes Routes) {

	srv := &http.Server{
		ReadTimeout:  200 * time.Second,
		WriteTimeout: 20000 * time.Second,
		Addr:         ":8080",
		Handler:      NewRouter(routes),
	}
	fmt.Println("listen")
	log.Fatal(srv.ListenAndServe())
}
