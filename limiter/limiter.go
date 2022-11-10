package limiter

import (
	"context"
	"strconv"
	"strings"
	"time"
)

type (
	Limiter interface {
		AllowN(ctx context.Context, key string, limit Limit, n int) (*Result, error)
	}
	Limit struct {
		Rate   int
		Burst  int
		Period time.Duration
	}
	Result struct {
		Limit     Limit
		Allowed   bool  // 是否允许访问
		Remaining int   // 剩余token数量
		ResetAt   int64 // bucket下次重置大致的(不准确)时间戳(unix timestamp)
	}
)

// Parse a string to Limit, the string MUST be formatted like <rate>,<burst>,<duration>
func Parse(s string) (l Limit) {
	l = PerSecond(1, 1)
	sections := strings.Split(s, ",")
	if len(sections) != 3 {
		return
	}

	rate, err := strconv.Atoi(sections[0])
	if err != nil {
		return
	}
	burst, err := strconv.Atoi(sections[1])
	if err != nil {
		return
	}
	period, err := time.ParseDuration(sections[2])
	if err != nil {
		return
	}

	return Every(rate, burst, period)
}

func PerSecond(rate, burst int) Limit {
	return Every(rate, burst, time.Second)
}

func PerMinute(rate int) Limit {
	return Every(rate, rate, time.Minute)
}

func Every(rate, burst int, period time.Duration) Limit {
	if burst < rate {
		burst = rate
	}
	return Limit{
		Rate:   rate,
		Burst:  burst,
		Period: period,
	}
}

var (
	TimeNow = func() time.Time { return time.Now() }
)
