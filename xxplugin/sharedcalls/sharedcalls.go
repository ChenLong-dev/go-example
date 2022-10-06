package sharedcalls

import (
	"sync"
)

// SharedCalls lets the concurrent calls with same key share the same result
type SharedCalls interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
}

// NewSharedCalls creates a new SharedCalls
func NewSharedCalls() SharedCalls {
	return &sharedGroup{
		calls: make(map[string]*call),
	}
}

type (
	sharedGroup struct {
		calls map[string]*call
		sync.Mutex
	}

	call struct {
		wg sync.WaitGroup

		val interface{}
		err error
	}
)

func (sg *sharedGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	c, done := sg.createCall(key)
	if done {
		return c.val, c.err
	}

	sg.makeCall(c, key, fn)
	return c.val, c.err
}

func (sg *sharedGroup) createCall(key string) (*call, bool) {
	sg.Lock()
	if c, ok := sg.calls[key]; ok {
		sg.Unlock()
		c.wg.Wait()
		return c, ok
	}

	c := new(call)
	c.wg.Add(1)
	sg.calls[key] = c
	sg.Unlock()

	return c, false
}

func (sg *sharedGroup) makeCall(c *call, key string, fn func() (interface{}, error)) {
	defer func() {
		sg.Lock()
		c.wg.Done()
		delete(sg.calls, key)
		sg.Unlock()
	}()

	c.val, c.err = fn()
}
