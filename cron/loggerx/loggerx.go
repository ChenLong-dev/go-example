package loggerx

import (
	"github.com/zeromicro/go-zero/core/logx"
)

var DefaultLogger = &CronLogger{}

type CronLogger struct{}

func (l *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	logx.Infow(msg, sweetenFields(keysAndValues)...)
}

func (l *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	fields := make([]logx.LogField, 0, len(keysAndValues))
	fields = append(fields, logx.Field("error", err))
	logx.Errorw(msg, append(fields, sweetenFields(keysAndValues)...)...)
}

type invalidPair struct {
	position   int
	key, value interface{}
}

type invalidPairs []invalidPair

func sweetenFields(args []interface{}) []logx.LogField {
	if len(args) == 0 {
		return nil
	}

	// Allocate enough space for the worst case; if users pass only structured
	// fields, we shouldn't penalize them with extra allocations.
	fields := make([]logx.LogField, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(logx.LogField); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			logx.Errorw("Ignored key without a value.", logx.Field("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, logx.Field(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		logx.Infow("Ignored key-value pairs with non-string keys.", logx.Field("invalid", invalid))
	}
	return fields
}
