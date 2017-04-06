In this exercise we'll implement a simplified version of the [echo](https://en.wikipedia.org/wiki/Echo_(command)).

Try the followings in your shell:

```
$ echo hello
hello
```

```
$ echo hello world
hello world
```

## Instructions

Create `echo.go`, and complete the `main` function.

```
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range ??? {
		???
	}

	fmt.Print("\n")
}
```

+ [os.Args](https://golang.org/pkg/os/#pkg-variables)
+ [for ... range](https://golang.org/doc/effective_go.html#for)
+ [fmt.Print](https://golang.org/pkg/fmt/#Print)

## Examples

```
$ go run echo.go hello
hello
```

```
$ go run echo.go hello
hello world
```