// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the LICENSE file or <http://unlicense.org/>

// +build darwin linux netbsd openbsd solaris

package maxopen

import (
	"errors"
	"syscall"
	"testing"
)

func TestInitSucceeded(t *testing.T) {
	if Err() != nil {
		t.Fatalf("error during init(): %s", err)
	}
}

func TestInitValues(t *testing.T) {
	if initial == 0 {
		t.Fatal("error during init(): initial value is zero")
	}
	if current == 0 {
		t.Fatal("error during init(): current value is zero")
	}
	if maximum == 0 {
		t.Fatal("error during init(): maximum value is zero")
	}
}

func TestValueFunctions(t *testing.T) {
	if Initial() == 0 {
		t.Fatal("error during Initial(): zero value returned")
	}
	if Current() == 0 {
		t.Fatal("error during Current(): zero value returned")
	}
}

func TestSet(t *testing.T) {
	for i := current; i > 0; i-- {
		Set(i)
		if Err() != nil {
			t.Fatalf("error during Set(%d): %s", i, err)
		}
	}
}

func TestReset(t *testing.T) {
	Reset()
	if Err() != nil {
		t.Fatalf("error during Reset(): %s", err)
	}
	if current != initial {
		t.Fatalf("error during Reset(): current != initial (%d != %d)", current, initial)
	}
}

type getrlimitError struct{}

func (getrlimitError) Getrlimit(resource int, rlim *syscall.Rlimit) error {
	return errors.New("Forced error for testing purposes")
}

func (getrlimitError) Setrlimit(resource int, rlim *syscall.Rlimit) error {
	return syscall.Setrlimit(resource, rlim)
}

type setrlimitError struct{}

func (setrlimitError) Getrlimit(resource int, rlim *syscall.Rlimit) error {
	return syscall.Getrlimit(resource, rlim)
}

func (setrlimitError) Setrlimit(resource int, rlim *syscall.Rlimit) error {
	return errors.New("Forced error for testing purposes")
}

func TestErrors(t *testing.T) {
	savedErr := Err()
	savedCurrent := Current()
	savedRlimits := rlimits
	err = nil
	rlimits = getrlimitError{}
	Reset()
	if Err() == nil {
		t.Fatal("expected forced error")
	}
	if savedCurrent != Current() {
		t.Fatal("expected current to be unchanged")
	}
	err = nil
	rlimits = setrlimitError{}
	Reset()
	if Err() == nil {
		t.Fatalf("expected forced error")
	}
	if savedCurrent != Current() {
		t.Fatal("expected current to be unchanged")
	}
	err = savedErr
	rlimits = savedRlimits
}
