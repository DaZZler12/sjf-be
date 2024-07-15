package logger

import (
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/text/encoding/charmap"
)

var (
	logger *zap.Logger
	Once   sync.Once
)

// GetLoggerInstance returns a singleton instance of the logger
func GetLoggerInstance() *zap.Logger {
	Once.Do(func() {
		loggerConfig := zap.NewProductionEncoderConfig()
		loggerConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		writer := io.Writer(os.Stdout)
		logger = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(loggerConfig),
			zapcore.AddSync(charmap.ISO8859_1.NewEncoder().Writer(writer)),
			zapcore.DebugLevel,
		))
	})
	return logger
}
