// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the UNLICENSE file or <http://unlicense.org/>

// +build darwin linux netbsd openbsd solaris

package maxopen

import (
	"sync"
	"syscall"
)

var (
	initial uint64
	current uint64
	maximum uint64
	err     error
	mu      sync.Mutex
)

func init() {
	// get current value
	r := syscall.Rlimit{}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &r)
	if err != nil {
		return
	}

	// store initial value
	// save current at this point in case an error occurs before completing
	initial = r.Cur
	current = r.Cur
	maximum = r.Max

	// attempt to set to maximum
	current, err = set(r.Max)

}

// Err returns any error that occured while updating the resource limit.
func Err() error {
	mu.Lock()
	defer mu.Unlock()
	return err
}

// Initial returns the original value of the resource limit.
// It is not modified.
func Initial() uint64 {
	return initial
}

// Current returns the most-recently observed value of the resource limit.
// It is modified by successful calls to Set() or Reset()
func Current() uint64 {
	mu.Lock()
	defer mu.Unlock()
	return current
}

// Set attempts to change the current limit to the specified value.
func Set(n uint64) {
	mu.Lock()
	defer mu.Unlock()
	current, err = set(n)
}

// Set attempts to change the current limit to the specified value.
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	current, err = set(initial)
}

func set(n uint64) (uint64, error) {
	r := syscall.Rlimit{Cur: n, Max: maximum}
	e := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &r)
	if e != nil {
		return current, e
	}
	e = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &r)
	if e != nil {
		return current, e
	}
	return r.Cur, err
}
