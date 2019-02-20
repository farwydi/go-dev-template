package dev_tempalte_logger

import (
	"github.com/farwydi/go-dev-template"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strings"
)

func InitZapWithLumberjack(logFile string) *zap.Logger {
	return zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	))
}

func InitZapStdOutDebug() *zap.Logger {
	return zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	))
}

func SQLError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	fns := strings.Split(fn, "/")
	dev_tempalte.ZapLogger.Error("SQL Error",
		zap.String("func", strings.Join([]string{fns[len(fns)-2], fns[len(fns)-1]}, "/")),
		zap.Int("line", line),
		zap.Error(err),
	)
}

func SQLPanic(err error) {
	_, fn, line, _ := runtime.Caller(1)
	fns := strings.Split(fn, "/")
	dev_tempalte.ZapLogger.Panic("SQL Panic",
		zap.String("func", strings.Join([]string{fns[len(fns)-2], fns[len(fns)-1]}, "/")),
		zap.Int("line", line),
		zap.Error(err),
	)
}
