package tasks

import (
	"github.com/bartdeboer/script/v2"
)

// Run tests using bitfield/script
func Test() {
	script.Exec("go test . -count=1 -v").Stdout()
}
