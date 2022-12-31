package router

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
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

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Content-Type") != "application/json" {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusUnsupportedMediaType)
				response, _ := json.Marshal(map[string]string{
					"error": "Incorrect content type. Expected application/json",
				})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func authenticationMiddleware(db *gorm.DB) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

//func panicHandlerMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer func() {
//			if err := recover(); err != nil {
//				w.WriteHeader(http.StatusInternalServerError)
//				log.Fatalln(err)
//			}
//		}()
//		next.ServeHTTP(w, r)
//	})
//}
