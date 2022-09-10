package logger

import (
	"testing"

	"go.uber.org/zap"
)

// go test -v
func TestDebug(t *testing.T) {

	lfg := new(LogConfig)
	lfg.Filename = "/tmp/xxx"
	//WithMultiFile(lfg)
	logger := NewLogger(lfg)

	logger.GetZlog().Debug("Debug", zap.Int("int", 00))
	logger.GetZlog().Info("Info", zap.Int("int", 12))
	logger.GetZlog().Warn("Warn", zap.Int("int", 13))
	logger.GetZlog().Debug("Debug", zap.Int("int", 14))
	logger.GetZlog().Sugar().Warn("Debug", zap.Int("int", 14))

	lfg2 := new(LogConfig)
	lfg2.Filename = "/tmp/xxx2222"
	WithMultiFile(lfg2)
	logger2 := NewLogger(lfg2)

	logger2.GetZlog().Debug("Debug", zap.Int("int", 00))
	logger2.GetZlog().Info("Info", zap.Int("int", 12))
	logger2.GetZlog().Warn("Warn", zap.Int("int", 13))
	logger2.GetZlog().Debug("Debug", zap.Int("int", 14))
	logger2.GetZlog().Sugar().Warn("Debug", zap.Int("int", 14))
	logger2.GetZlog().Error("error")
}

// go test -bench=. -v
func BenchmarkDebug(b *testing.B) {
	lfg := new(LogConfig)
	lfg.Filename = "/tmp/test2.log"
	logger := NewLogger(lfg)
	for i := 1; i < b.N; i = i + 1 {
		logger.GetZlog().Info("abc", zap.Int("int", 11))
	}
}
