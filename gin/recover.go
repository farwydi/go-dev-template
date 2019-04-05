package dev_tempalte_gin

import (
    "github.com/farwydi/go-dev-template"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "net"
    "net/http"
    "os"
    "strings"
)

func Recovery(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            // Check for a broken connection, as it is not really a
            // condition that warrants a panic stack trace.
            var brokenPipe bool
            if ne, ok := err.(*net.OpError); ok {
                if se, ok := ne.Err.(*os.SyscallError); ok {
                    if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
                        brokenPipe = true
                    }
                }
            }

            path := c.Request.URL.Path
            query := c.Request.URL.RawQuery

            dev_tempalte.ZapLogger.Info("Recover panic",
                zap.String("method", c.Request.Method),
                zap.String("path", path),
                zap.String("query", query),
                zap.Reflect("err", err),
            )

            // If the connection is dead, we can't write a status to it.
            if brokenPipe {
                c.Error(err.(error)) // nolint: errcheck
                c.Abort()
            } else {
                c.AbortWithStatus(http.StatusInternalServerError)
            }
        }
    }()

    c.Next()
}
