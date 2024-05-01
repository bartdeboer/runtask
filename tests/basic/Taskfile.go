package basic

import (
	"fmt"
	"strings"
)

// A simple test
func Simple() {
	fmt.Println("Hello, World!")
}

// A simple argument test
func SimpleArg(arg string) {
	fmt.Printf("Hello, %s!\n", arg)
}

// Variadic arguments
func Variadic(arg1 string, args ...string) {
	fmt.Printf("Hello, %s [%s]!\n", arg1, strings.Join(args, ", "))
}

// Addition #1
func Add1(arg1 float64, int1, int2, int3, int4, int5 int) {
	fmt.Printf("%.3f %d\n", arg1, int1+int2+int3+int4+int5)
}

// Addition #2
func Add2(arg1 float64, args ...int) {
	total := 0
	for _, arg := range args {
		total += arg
	}
	fmt.Printf("%.3f %d\n", arg1, total)
}

// Override run function
func run() {
	fmt.Println("Overridden")
}

// Run task
func Run() {
	run()
}
