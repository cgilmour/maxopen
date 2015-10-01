// Copyright 2015 Caleb Gilmour
// Use of this source code free and unencumbered software released into the public domain.
// For more information, refer to the UNLICENSE file or <http://unlicense.org/>

package main

import (
	_ "github.com/cgilmour/maxopen"
	"fmt"
)

func main() {
	fmt.Println("Imported github.com/cgilmour/maxopen  only for side-effects: initializing max open file limit to the maximum permitted")
	fmt.Println("Now go wild with your http server or whatnot. No need to wrap it in a script that runs ulimit -n")
}

