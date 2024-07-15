package lock

import (
	"fmt"
	"os"
	"syscall"
)

type Flock struct {
	file string
	fp   *os.File
}

func (fl *Flock) Lock() error {
	if err := syscall.Flock(int(fl.fp.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		return err
	}

	return nil
}

func (fl *Flock) Unlock() {
	if err := syscall.Flock(int(fl.fp.Fd()), syscall.LOCK_UN); err != nil {
		fmt.Errorf("syscall.Flock Unlock fail: %v", err)
	}
}

func (fl *Flock) Close() {
	fl.Unlock()
	fl.fp.Close()
}

func NewFileLock(path string) (*Flock, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("%s is not exist! err: %v", path, err)
	}

	if info.IsDir() {
		dp, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("OpenDir err: %v", err)
		}
		return &Flock{path, dp}, nil
	} else {
		fp, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, fmt.Errorf("OpenFile err: %v", err)
		}

		return &Flock{path, fp}, nil
	}
}
