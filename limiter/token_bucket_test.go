package limiter

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	red "github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func setupRedis() (s *miniredis.Miniredis, rdb red.UniversalClient, clean func()) {
	s, _ = miniredis.Run()
	rdb = red.NewUniversalClient(&red.UniversalOptions{
		Addrs:        []string{s.Addr()},
		DialTimeout:  1 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})

	clean = func() {
		rdb.Close()
		s.Close()
	}
	return
}

func TestTokenBucket_Allow(t *testing.T) {
	Convey("rate limiter(token bucket), 测试rate和burst\n", t, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := TokenBucket(rdb)
		rate, burst := 10, 15
		lim := PerSecond(rate, burst)
		requested := 1
		api := "GET:/admin/api/v1/users"

		TimeNow = func() func() time.Time {
			n := time.Now()
			return func() time.Time {
				return n
			}
		}()
		Convey(fmt.Sprintf("rate limiter(token bucket): 每秒允许(rate=%d, burst=%d)个请求\n", rate, burst),
			func() {
				for i := 1; i <= rate; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					//t.Logf("result:%v", result)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeTrue)
				}
			})
		Convey(fmt.Sprintf("rate limiter(token bucket): 允许(rate=%d, burst=%d)个突增请求\n", rate, burst),
			func() {
				for i := 1; i <= burst; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					//t.Logf("result:%v", result)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeTrue)
				}
			})
		Convey(fmt.Sprintf("rate limiter(token bucket): 超出限额(rate=%d, busrt=%d)的访问请求会失败\n", rate, burst),
			func() {
				for i := 1; i <= burst; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					//t.Logf("result1:%v", result)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeTrue)
				}
				for i := 1; i <= 10; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					//t.Logf("result2:%v", result)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeFalse)
				}
			})
	})
}

func TestTokenBucket_AllowForDifferentReq(t *testing.T) {
	Convey("rate limiter(token bucket): 不同请求的访问频率会隔离", t, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := TokenBucket(rdb)
		rate, burst := 10, 15
		lim := PerSecond(rate, burst)
		requested := 1
		api1, api2 := "GET:/admin/api/v1/users", "POST:/admin/api/v1/users"

		TimeNow = func() func() time.Time {
			n := time.Now()
			return func() time.Time {
				return n
			}
		}()

		// allowed request
		for i := 1; i <= burst; i++ {
			ctx := context.Background()
			result, err := limiter.AllowN(ctx, api1, lim, requested)
			//t.Logf("result api1:%v", result)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeTrue)

			result, err = limiter.AllowN(ctx, api2, lim, requested)
			//t.Logf("result api2:%v", result)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeTrue)
		}

		// forbidden request
		for i := 1; i <= 10; i++ {
			ctx := context.Background()
			result, err := limiter.AllowN(ctx, api1, lim, requested)
			//t.Logf("result api1 10:%v", result)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeFalse)

			result, err = limiter.AllowN(ctx, api2, lim, requested)
			//t.Logf("result api2 10:%v", result)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeFalse)
		}
	})
}

func TestTokenBucket_ForHugeConcurrentRequest(t *testing.T) {
	Convey("rate limiter(token bucket), 超出允许范围内的(rate=%d, burst=%d), 大流量并发请求会失败", t, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := TokenBucket(rdb)
		rate, burst := 10, 15
		lim := PerSecond(rate, burst)
		requested := burst * 2
		//requested := 5
		api := "GET:/admin/api/v1/users"

		TimeNow = func() func() time.Time {
			n := time.Now()
			return func() time.Time {
				return n
			}
		}()

		for i := 1; i <= 10; i++ {
			ctx := context.Background()
			result, err := limiter.AllowN(ctx, api, lim, requested)
			//t.Logf("result:%v", result)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeFalse)
		}
	})
}

func TestToken_BucketRefill(t *testing.T) {
	Convey("rate limiter(token bucket), token会被持续refill", t, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := TokenBucket(rdb)
		rate, burst := 10, 15
		lim := PerSecond(rate, burst)
		requested := 1
		api := "GET:/admin/api/v1/users"

		n := time.Now()
		TimeNow = func() time.Time {
			return n
		}
		for i := 1; i <= burst; i++ {
			ctx := context.Background()
			result, err := limiter.AllowN(ctx, api, lim, requested)
			//t.Logf("result1:%v", result)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeTrue)
		}

		// later
		for i := 1; i <= 10; i++ {
			later := n.Add(time.Duration(i) * time.Second)
			TimeNow = func() time.Time {
				return later
			}

			for i := 1; i <= rate; i++ {
				ctx := context.Background()
				result, err := limiter.AllowN(ctx, api, lim, requested)
				//t.Logf("result2:%v", result)
				So(err, ShouldBeNil)
				So(result.Allowed, ShouldBeTrue)
			}
		}
	})
}

func TestTokenBucket_ReturnErrOnAnyRedisErr(t *testing.T) {
	Convey("rate limiter(token bucket), 当redis出现error时, AllowN返回error", t, func() {
		s, rdb, _ := setupRedis()
		defer rdb.Close()

		limiter := TokenBucket(rdb)
		rate, burst := 10, 15
		lim := PerSecond(rate, burst)
		requested := 1
		api := "GET:/admin/api/v1/users"

		// close redis
		s.Close()
		ctx := context.Background()
		result, err := limiter.AllowN(ctx, api, lim, requested)
		t.Logf("result:%v", result)
		So(err, ShouldNotBeNil)
		So(result, ShouldBeNil)
	})
}

// go test --bench=TokenBucket token_bucket_test.go -benchmem
func BenchmarkTokenBucket(b *testing.B) {
	Convey("token bucket benchmark", b, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := TokenBucket(rdb)
		rate, burst := 10, 15
		lim := PerSecond(rate, burst)
		requested := 1
		api := "GET:/admin/api/v1/users"

		for n := 0; n < b.N; n++ {
			ctx := context.Background()
			_, err := limiter.AllowN(ctx, api, lim, requested)
			So(err, ShouldBeNil)
		}
	})
}
