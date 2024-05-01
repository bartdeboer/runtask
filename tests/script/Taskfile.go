package script

import (
	"fmt"
	"io"
	"strings"

	"github.com/bitfield/script"
)

// Script file test
func Count() {
	lines, _ := script.File("../../testdata/test.txt").CountLines()
	fmt.Println(lines)
}

// Script to upper test
func ToUpper() {
	script.File("../../testdata/test.txt").FilterLine(strings.ToUpper).Stdout()
}

func Ping(addr string) {
	script.Exec(fmt.Sprintf("ping %s", addr)).Stdout()
}

// Exec test
func ExecGoVersion() {
	script.Exec("go version").Stdout()
}

// Text filter test
func TextFilter() {
	script.File("../../testdata/test.txt").FilterScan(func(line string, w io.Writer) {
		fmt.Fprintf(w, "scanned line: %q\n", line)
	}).Stdout()
}
