package limiter

import (
	red "github.com/go-redis/redis/v8"
)

type rediser interface {
	red.Scripter
}
