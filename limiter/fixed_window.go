package limiter

import (
	"context"
	red "github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// FixedWindow 基于redis的fixed window算法
func FixedWindow(rdb rediser) Limiter {
	return &fw{rdb: rdb}
}

type fw struct {
	rdb rediser
}

const FwkeyPrefix = "rate:fw:"

func fwkey(k string) string {
	return FwkeyPrefix + k
}

func (l *fw) AllowN(ctx context.Context, key string, lim Limit, n int) (*Result, error) {
	keys := []string{fwkey(key)}
	args := []interface{}{
		strconv.Itoa(lim.Rate),
		lim.Period.Seconds(),
		strconv.Itoa(n),
	}

	v, err := fwscript.Run(ctx, l.rdb, keys, args...).Result()
	if err != nil {
		return nil, err
	}

	values := v.([]interface{})
	remainingTTL := values[2].(int64)
	return &Result{
		Limit:     lim,
		Allowed:   values[0].(string) == "true",
		Remaining: int(values[1].(int64)),
		ResetAt:   TimeNow().Add(time.Duration(remainingTTL) * time.Second).UTC().Unix(),
	}, nil
}

// refer to go-zero
var fwscript = red.NewScript(`
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local requested = tonumber(ARGV[3])
local current = redis.call("INCRBY", KEYS[1], requested)

local ttl = window
if current > requested then
	ttl = redis.call("ttl", KEYS[1])
end

if current == requested then
    redis.call("expire", KEYS[1], window)
    return {
		tostring(limit >= current),
		limit - current,
		ttl -- remaining ttl
	} 
elseif current <= limit then
    return {
		"true",
		limit - current,
		ttl,
	} 
else
    return {
		"false",
		0,
		ttl
	} 
end
`)
