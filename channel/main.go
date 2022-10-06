package main

import (
	"fmt"
	"time"
)

// 超时控制
func do() <-chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		// TODO:do something
		time.Sleep(6 * time.Second)
		done <- struct{}{}
	}()
	return done
}

func test01() {
	done := do()
	select {
	case <-done:
		fmt.Println("is done ...")
	case <-time.After(5 * time.Second):
		fmt.Println("is timeout ...")
	default:
		fmt.Println("no black ...")
	}
}

// 取最快的结果
func call(ret chan<- string, n int) {
	// TODO:do samething

	ret <- fmt.Sprintf("result-%d", n)
}

func test02() {
	ret := make(chan string, 3)
	for i := 0; i < cap(ret); i++ {
		go func(n int) {
			call(ret, n)
		}(i)
	}
	fmt.Println(<-ret)
}

//  限制最大并发数
func test03() {
	// 最大并发数为 3
	limits := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		go func(n int) {
			// 缓冲区满了就会阻塞在这
			limits <- struct{}{}
			fmt.Printf("n:%d, chan len:%d\n", n, len(limits))
			do2(n)
			<-limits
		}(i)
		fmt.Printf("i:%d, chan len:%d\n", i, len(limits))
	}
	select {}
}

func do2(n int) {
	fmt.Printf("[do2] n:%d\n", n)
	//time.Sleep(3 * time.Second)
	for {
		time.Sleep(3 * time.Second)
		fmt.Printf("===[do2] n:%d\n", n)
	}
}

//for...range 优先
func test04() {
	c := make(chan int, 20)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	fmt.Printf("1 chan len:%d\n", len(c))
	// 当 c 被关闭后，取完里面的元素就会跳出循环
	for x := range c {
		fmt.Println(x)
	}
	fmt.Printf("2 chan len:%d\n", len(c))
}

//多个 goroutine 同步响应
func do5(c <-chan struct{}, n int) {
	// 会阻塞直到收到 close
	<-c
	fmt.Printf("hello-n:%d\n", n)
}

func test05() {
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(n int) {
			do5(c, n)
		}(i)
	}
	close(c)
}

func main() {
	//test01()
	//test02()
	//test03()
	//test04()
	test05()

}
