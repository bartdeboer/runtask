package env

import (
	"fmt"
	"os"
)

// Environment test
func Simple() {
	os.Setenv("BAR", "Hello, John!")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
}
