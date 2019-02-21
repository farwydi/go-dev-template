package dev_tempalte_http

import (
	"github.com/farwydi/go-dev-template"
	"go.uber.org/zap"
	"net/http"
)

type recoveryHandler struct {
	handler    http.Handler
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			dev_tempalte.ZapLogger.Error("Recovery", zap.Reflect("error", err))
		}
	}()

	h.handler.ServeHTTP(w, r)
}

func RecoveryHandler(h http.Handler) http.Handler {
	return recoveryHandler{h}
}
