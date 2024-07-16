package reslimits

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

type Pool struct {
	wg           sync.WaitGroup
	queue        chan struct{}
	maxGoroutine int
	defaultNum   int
}

var gPool *Pool

// NewGoPool 实例化一个go程池
func NewGoPool(maxGoroutine int) (err error) {
	gPool = &Pool{
		queue:        make(chan struct{}, maxGoroutine),
		maxGoroutine: maxGoroutine,
		defaultNum:   runtime.NumGoroutine(),
	}
	fmt.Printf("NewGoPool is successful! gonum:%d\n", maxGoroutine)
	return
}

func (p *Pool) add() {
	if p.maxGoroutine == 0 {
		return
	}
	p.queue <- struct{}{}
	p.wg.Add(1)
	fmt.Printf("+ add a goroutine:%d\n", GetGid())
}

func (p *Pool) done() {
	if p.maxGoroutine == 0 {
		return
	}
	p.wg.Done()
	<-p.queue
	fmt.Printf("- del a goroutine:%d\n", GetGid())
}

func (p *Pool) wait() {
	if p.maxGoroutine == 0 {
		return
	}
	p.wg.Wait()
	fmt.Printf("= wait all goroutine\n")
}

func (p *Pool) getNum() int {
	return p.maxGoroutine
}

func (p *Pool) getCurrentNum() int {
	return runtime.NumGoroutine() - p.defaultNum
}

// Add 添加
func Add() {
	gPool.add()
}

// Done 释放
func Done() {
	gPool.done()
}

// Wait 等待
func Wait() {
	gPool.wait()
}

// GetDefaultNum 获取默认协程数量
func GetDefaultNum() int {
	return gPool.defaultNum
}

// GetNum 获取设置的协程数量
func GetNum() int {
	return gPool.getNum()
}

// GetCurrentNum 获取当前的协程数量
func GetCurrentNum() int {
	return gPool.getCurrentNum()
}

func GetGid() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0
	}
	return n
}
