package benchmarktest

import "testing"

/*
$ go mod init benchmarktest
$ go mod tidy

$ go test -bench=. -run=none
goos: windows
goarch: amd64
pkg: benchmarktest
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_fib-8          4510855               266.7 ns/op
PASS
ok      benchmarktest   3.264s

$ go test -bench=. -run=Benchmark_fib
goos: windows
goarch: amd64
pkg: benchmarktest
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_fib-8          4588830               263.0 ns/op
PASS
ok      benchmarktest   3.494s

$ go test -bench=fib_test.go -run=Benchmark_fib
PASS
ok      benchmarktest   1.661s


-v 				# 打印详细日志

-bench=.		# 运行所有的基准测试
-bench=_fib		# 只运行Benchmark_fib
-bench=fib$		# 匹配所有fib结尾的
-benchtime=3s	# 每一轮运行3秒
-benchtime=300x	# 循环运行300次
-cpu=2,4		# 指定GOMAXPROCS的数量，模拟多核。分别2核和4核运行一次测试
-count=3		# 运行3轮

-benchmem		# 显示堆内存分配情况，分配的越多越影响性能
*/
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func Benchmark_fib(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(10)
	}
}
