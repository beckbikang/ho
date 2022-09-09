package logger

import (
	"testing"

	"go.uber.org/zap"
)

// go test -v
func TestDebug(t *testing.T) {

	lfg := new(LogConfig)
	lfg.Filename = "/tmp/test.log"
	logger := NewLogger(lfg)

	logger.GetZlog().Info("test")
	logger.GetZlog().Info("abc", zap.Int("int", 11))

	logger.GetZlog().Info("test122131")

	logger.zlogs.Sugar().Info("ttttt")
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
