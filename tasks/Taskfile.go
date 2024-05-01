package tasks

import (
	"github.com/bitfield/script"
)

// Run tests using bitfield/script
func Test() {
	script.Exec("go test . -count=1 -v").Stdout()
}
