package dev_tempalte_universal

import (
	"github.com/farwydi/go-dev-template"
	"go.uber.org/zap"
	"time"
)

// Restart after panic, timeout 5 sec
// With message
func Restart(f func(), msg string) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case error:
				dev_tempalte.ZapLogger.Info("Recover panic", zap.Error(t))
			case string:
				dev_tempalte.ZapLogger.Info("Recover panic", zap.String("error", t))
			}
		}

		time.Sleep(time.Second * 5)
		go Restart(f, msg)
	}()

	dev_tempalte.ZapLogger.Info(msg)

	f()
}