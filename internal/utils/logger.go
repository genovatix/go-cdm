package utils

import (
	"go.uber.org/zap"
)

func GetLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}
