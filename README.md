Runtask is a task runner that enables writing Taskfile scripts in Go.

## Install & Usage

Install runtask globally:

```
go install github.com/bartdeboer/runtask@latest
```

and start writing your Taskfiles anywhere. Any capitalized function should be
available to run from the command line:

```go
import (
	"fmt"
	"io"
	"strings"

	"github.com/bitfield/script"
)

// A simple argument test
func SimpleArg(arg string) {
	fmt.Printf("Hello, %s!\n", arg)
}

// Variadic arguments
func Variadic(arg1 string, args ...string) {
	fmt.Printf("Hello, %s [%s]!\n", arg1, strings.Join(args, ", "))
}

// Ping some address
func Ping(addr string) {
	script.Exec(fmt.Sprintf("ping %s", addr)).Stdout()
}
```

And run your tasks:

```
runtask simplearg World
runtask variadic John Jane Johnie
runtask ping 127.0.0.1
runtask help variadic
```

Runtask looks for the following files:

* `Taskfile`
* `Taskfile.go`
* `tasks/Taskfile.go`

A package name is optional but can be helpful for enabling language support features

Runtask uses [traefik/yaegi](https://github.com/traefik/yaegi) as its Go interpreter.

Next to Go's standard library, it also includes [bitfield/script](https://github.com/bitfield/script)
for convenience.
