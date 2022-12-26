package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type loggedResponse struct {
	http.ResponseWriter
	statusCode int
}

func (l *loggedResponse) WriteHeader(status int) {
	l.statusCode = status
	l.ResponseWriter.WriteHeader(status)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggedResp := &loggedResponse{ResponseWriter: w}
		start := time.Now()
		next.ServeHTTP(loggedResp, r)
		log.Println(fmt.Sprintf("[%s]: '%s' %d - %s", r.Method, r.URL.Path, loggedResp.statusCode, time.Since(start)))
	})
}

func panicHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalln(err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
