package dev_tempalte_http

import (
	"github.com/farwydi/go-dev-template"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

type loggingHandler struct {
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	sw := statusWriter{ResponseWriter: w}
	h.handler.ServeHTTP(&sw, r)

	var query []zap.Field
	query = append(query, zap.Duration("exec", time.Since(t)))
	query = append(query, zap.String("method", r.URL.Path))
	query = append(query, zap.Int("status", sw.status))
	query = append(query, zap.String("ip", r.RemoteAddr))
	query = append(query, zap.String("xff", r.Header.Get("X-Forwarded-For")))
	query = append(query, zap.Reflect("query", r.URL.Query()))
	dev_tempalte.ZapLogger.Info("Access", query...)
}

func LoggingHandler(h http.Handler) http.Handler {
	return loggingHandler{h}
}
