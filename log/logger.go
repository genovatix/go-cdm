package log

import (
	"go.uber.org/zap"
)

// Logger is the shared logger instance.
var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
