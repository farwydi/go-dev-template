package dev_tempalte_path

import (
	"github.com/farwydi/go-dev-template"
	"go.uber.org/zap"
	"os"
)

// Make directory
// panic
func MakeDirWithPanic(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dev_tempalte.ZapLogger.Info("Make folder", zap.String("dir", dir))
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			dev_tempalte.ZapLogger.Panic("Fail make dir", zap.Error(err))
		}
	}
}

func MakeDirSafe(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dev_tempalte.ZapLogger.Info("Make folder", zap.String("dir", dir))
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
