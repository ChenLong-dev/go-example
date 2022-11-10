package limiter

import (
	"context"
	"strconv"

	red "github.com/go-redis/redis/v8"
)

// TokenBucket 基于redis的token bucket算法
// 注意: 该实现只支持基于秒的(req/sec)访问限制, 即AllowN()函数中的limit.Limit.Period只能为1秒
func TokenBucket(rdb rediser) Limiter {
	return &tb{rdb: rdb}
}

type tb struct {
	rdb rediser
}

const tbkeyPrefix = "rate:tb:"

func tbkey(k string) string {
	return tbkeyPrefix + k
}

func (l *tb) AllowN(ctx context.Context, key string, lim Limit, n int) (*Result, error) {
	now := TimeNow()
	keys := []string{tbkey(key)}
	args := []interface{}{
		strconv.Itoa(lim.Rate),
		strconv.Itoa(lim.Burst),
		strconv.FormatInt(now.Unix(), 10),
		strconv.Itoa(n),
	}

	v, err := tbscript.Run(ctx, l.rdb, keys, args...).Result()
	if err != nil {
		return nil, err
	}

	values := v.([]interface{})
	return &Result{
		Limit:     lim,
		Allowed:   values[0].(string) == "true",
		Remaining: int(values[1].(int64)),
		ResetAt:   values[2].(int64),
	}, nil
}

// refer to go-zero
var tbscript = red.NewScript(`
local rate = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])

local last = redis.call("hmget", KEYS[1], "tokens", "timestamp")
local last_tokens = tonumber(last[1])
local first_call = false
if last_tokens == nil then
    last_tokens = capacity
	first_call = true
end

local last_refreshed = tonumber(last[2])
if last_refreshed == nil then
    last_refreshed = 0
end

local filled_tokens = math.min(capacity, last_tokens)
if not first_call then
	local delta = math.max(0, now-last_refreshed)
	filled_tokens = math.min(capacity, last_tokens+(delta*rate))
end

local allowed = filled_tokens >= requested
local new_tokens = filled_tokens
if allowed then
    new_tokens = filled_tokens - requested
end

redis.call("hset", KEYS[1], "tokens", new_tokens, "timestamp", now)
if first_call then
	local fill_time = capacity/rate
	local ttl = math.floor(fill_time*2)
	redis.call("expire", KEYS[1], ttl)
end

return {
    tostring(allowed), -- true or false
    new_tokens, -- remaining
    now + 1, -- reset timestamp
}
`)
