# maxopen
Automatically increase the Max Open Files limit for your go process.

Instead of wrapping your Go program in a script or launch environment settings, simply import this package.
You can import only for side-effects using underscore-import, or add a logging message using details captured during initialization.

Simple example:
```go
package main

import (
	_ "github.com/cgilmour/maxopen"
	"fmt"
)

func main() {
	fmt.Println("Imported github.com/cgilmour/maxopen  only for side-effects: initializing max open file limit to the maximum permitted")
	fmt.Println("Now go wild with your http server or whatnot. No need to wrap it in a script that runs ulimit -n")
}
```
