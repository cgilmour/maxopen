// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the LICENSE file or <http://unlicense.org/>

package main

import (
	"fmt"
	"github.com/cgilmour/maxopen"
)

func main() {
	fmt.Println("Used github.com/cgilmour/maxopen package")
	if maxopen.Err() != nil {
		fmt.Println("However, an error occurred. Your program should continue normally.")
		return
	}
	fmt.Println("Increased max open file limit from", maxopen.Initial(), "to", maxopen.Current())
	fmt.Println("Calling Reset()")
	maxopen.Reset()
	fmt.Println("Current limit is now", maxopen.Current())
	if maxopen.Err() != nil {
		fmt.Println("error occurred:", maxopen.Err())
	}
}
