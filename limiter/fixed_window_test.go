package limiter

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFixedWindow_Allow(t *testing.T) {
	Convey("rate limiter(fiexd window), 测试rate", t, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := FixedWindow(rdb)
		rate := 10
		lim := PerMinute(rate)
		requested := 1
		api := "POST:/sso/api/v1/123/sms"
		Convey(fmt.Sprintf("rate limiter(fixed window): 每分钟允许(rate=%d)个请求\n", rate),
			func() {
				for i := 1; i <= rate; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					//t.Logf("result:%v", result)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeTrue)
					So(result.Remaining, ShouldEqual, rate-i)
				}
			})
		Convey(fmt.Sprintf("rate limiter(token bucket): 超出限额(rate=%d)的访问请求会失败\n", rate),
			func() {
				// 消耗额度
				for i := 1; i <= rate; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeTrue)
					So(result.Remaining, ShouldEqual, rate-i)
				}
				// 限额请求
				for i := 1; i <= 10; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeFalse)
					So(result.Remaining, ShouldEqual, 0)
				}
			})
		Convey("rate limiter(fixed window): token会自动refill\n",
			func() {
				for i := 1; i <= rate; i++ {
					ctx := context.Background()
					result, err := limiter.AllowN(ctx, api, lim, requested)
					So(err, ShouldBeNil)
					So(result.Allowed, ShouldBeTrue)
				}

				// expire the counter
				for i := 0; i < 10; i++ {
					key := fwkey(api)
					rdb.Expire(context.Background(), key, 0)
					//time.Sleep(1 * time.Second)
					for i := 1; i <= rate; i++ {
						ctx := context.Background()
						result, err := limiter.AllowN(ctx, api, lim, requested)
						So(err, ShouldBeNil)
						So(result.Allowed, ShouldBeTrue)
						So(result.Remaining, ShouldEqual, rate-i)
					}
					for i := 1; i <= 10; i++ {
						ctx := context.Background()
						result, err := limiter.AllowN(ctx, api, lim, requested)
						So(err, ShouldBeNil)
						So(result.Allowed, ShouldBeFalse)
					}
				}
			})
	})
}

func TestFixedWindow_AllowForDifferentReq(t *testing.T) {
	Convey("rate limiter(fixed window): 不同请求的访问频率会隔离", t, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := FixedWindow(rdb)
		rate1, rate2 := 10, 20
		lim1, lim2 := PerMinute(rate1), PerMinute(rate2)
		requested := 1
		api1, api2 := "GET:/admin/api/v1/users", "POST:/admin/api/v1/users"

		// allowed request
		for i := 1; i <= rate1; i++ {
			ctx := context.Background()
			result, err := limiter.AllowN(ctx, api1, lim1, requested)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeTrue)
		}
		for i := 1; i <= rate2; i++ {
			ctx := context.Background()
			result, err := limiter.AllowN(ctx, api2, lim2, requested)
			So(err, ShouldBeNil)
			So(result.Allowed, ShouldBeTrue)
		}
	})
}

func TestFixedWindow_ReturnErrOnAnyRedisErr(t *testing.T) {
	Convey("rate limiter(fixed window), 当redis出现error时, AllowN返回error\n", t, func() {
		s, rdb, _ := setupRedis()
		defer rdb.Close()

		limiter := FixedWindow(rdb)
		rate := 10
		lim := PerMinute(rate)
		requested := 1
		api := "GET:/admin/api/v1/users"

		// close redis
		s.Close()
		ctx := context.Background()
		result, err := limiter.AllowN(ctx, api, lim, requested)
		So(err, ShouldNotBeNil)
		So(result, ShouldBeNil)
	})
}

func BenchmarkFixedWindow(b *testing.B) {
	Convey("fixed window benchmark", b, func() {
		_, rdb, clean := setupRedis()
		defer clean()

		limiter := FixedWindow(rdb)
		rate := 10
		lim := PerMinute(rate)
		requested := 1
		api := "GET:/admin/api/v1/users"

		for n := 0; n < b.N; n++ {
			ctx := context.Background()
			_, err := limiter.AllowN(ctx, api, lim, requested)
			So(err, ShouldBeNil)
		}
	})
}
