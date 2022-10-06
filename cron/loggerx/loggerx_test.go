package loggerx

import (
	"errors"
	"testing"
)

func TestLoggerx(t *testing.T) {
	DefaultLogger.Info("KEY", "key1", "123", "key2", []byte("xxxx"))

	err := errors.New("chenlong")
	DefaultLogger.Error(err, "KEY", "key1", 123, "key2", []byte("yyy"))
}
