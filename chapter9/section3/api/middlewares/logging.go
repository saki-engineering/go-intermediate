package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Println("res: ", rlw.code)
	})
}
