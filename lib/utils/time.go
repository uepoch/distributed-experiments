package utils

import (
	"time"

	"go.uber.org/zap"
)

func TimeTrack(logger *zap.Logger,t time.Time, message string) {
	logger.Debug(message, zap.Duration("duration", time.Since(t)))
}