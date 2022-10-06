package sharedcalls

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestExclusiveCallDoDupSuppress(t *testing.T) {
	g := NewSharedCalls()
	c := make(chan string)
	var calls int32
	fn := func() (interface{}, error) {
		atomic.AddInt32(&calls, 1)
		return <-c, nil
	}

	const n = 10
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			v, err := g.Do("key", fn)
			if err != nil {
				t.Errorf("Do error: %v", err)
			}
			if v.(string) != "bar" {
				t.Errorf("got %q; want %q", v, "bar")
			}
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond) // let goroutines above block
	c <- "bar"
	wg.Wait()
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Errorf("number of calls = %d; want 1", got)
	}
}

func TestExclusiveCallDoDupSuppress2(t *testing.T) {
	//g := NewSharedCalls()
	c := make(chan string)
	var calls int32
	fn := func() (interface{}, error) {
		calls = calls + 1
		return <-c, nil
	}

	const n = 10
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			//v, err := g.Do("key", fn)
			v, err := fn()
			if err != nil {
				t.Errorf("Do error: %v", err)
			}
			t.Logf("v:%v", v)
			//if v.(string) != "bar" {
			//	t.Errorf("got %q; want %q", v, "bar")
			//}
			wg.Done()
		}()
	}
	t.Logf("xxxxx01")
	time.Sleep(100 * time.Millisecond) // let goroutines above block
	c <- "bar"
	t.Log("calls:", calls)
	wg.Wait()

}
