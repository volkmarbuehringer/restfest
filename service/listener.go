package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/husobee/vestigo"
)

var logger = log.New(os.Stdout, "[req] ", 0)

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r.Method, "-", r.RequestURI)
		//Add data to context
		start := time.Now()
		ctx, cancel := context.WithCancel(r.Context())
		defer cancel() // releases resources if slowOperation completes before timeout elapses
		//	ctx := context.WithValue(r.Context(), "Username", "cookie.Value")

		go func() {
			time.Sleep(50 * time.Millisecond)
			cancel()
		}()
		next.ServeHTTP(w, r.WithContext(ctx))

		logger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func Listen(routes *vestigo.Router) {
	router := AddContext(routes)
	srv := &http.Server{
		ReadTimeout:  200 * time.Second,
		WriteTimeout: 20000 * time.Second,
		Addr:         ":8080",
		Handler:      router,
	}
	fmt.Println("listen")
	log.Fatal(srv.ListenAndServe())
}
