package t1

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

// Exec test
func Ping() {
	script.Exec("ping 127.0.0.1").Stdout()
}

func ExecGoVersion() {
	script.Exec("go version").Stdout()
}

// Text filter test
func TextFilter() {
	script.File("../../testdata/test.txt").FilterScan(func(line string, w io.Writer) {
		fmt.Fprintf(w, "scanned line: %q\n", line)
	}).Stdout()
}
