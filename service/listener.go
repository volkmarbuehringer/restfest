package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/husobee/vestigo"
	"github.com/sirupsen/logrus"
)

/*
type IRequestScopedLogger interface {
	Logger

	//GetRequestScoped(requestID, appName string, userID int64) Logger
}
*/

/*
type Logger interface {
	Printf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
*/
type IRequestScopedLogger struct {
	*logrus.Entry
	requestID string
	appName   string
	userID    int64
}

func NewLogger(requestID string, appName string, userID int64) *IRequestScopedLogger {
	l := logrus.New()
	e := l.WithFields(logrus.Fields{"request_id": requestID, "app_name": appName, "user_id": userID})

	return &IRequestScopedLogger{e, appName, requestID, userID}
}

/*
func (l *IRequestScopedLogger) Printf(format string, args ...interface{}) {
	l.WithFields(logrus.Fields{
		"datetime":   time.Now(),
		"request_id": l.requestID,
		"app_name":   l.appName,
		"user_id":    l.userID,
	}).Infof(format, args...)
}
*/
func NewContext(ctx context.Context, log *IRequestScopedLogger) context.Context {
	return context.WithValue(ctx, "logger", log)
}

func MustFromContext(ctx context.Context) *IRequestScopedLogger {
	u, ok := ctx.Value("logger").(*IRequestScopedLogger)
	if !ok {
		panic("user not found in context")
	}
	return u
}

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r.Method, "-", r.RequestURI)
		//Add data to context
		start := time.Now()

		reqlog := NewLogger(
			fmt.Sprintf("%x", 4711), // request id
			"web",    // application name
			int64(0), // current user id
		)

		lctx := NewContext(r.Context(), reqlog)
		r = r.WithContext(lctx)

		next.ServeHTTP(w, r.WithContext(lctx))

		reqlog.Infof("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
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
