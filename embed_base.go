package runtask

// Prevent conflicts
import (
	_fmt "fmt"
	_os "os"
	_exec "os/exec"
)

// Can be overridden from Taskfile
func run(command string, args ...string) {
	cmd := _exec.Command(command, args...)
	cmd.Stdout = _os.Stdout
	cmd.Stderr = _os.Stderr
	cmd.Stdin = _os.Stdin
	err := cmd.Run()
	if err != nil {
		_fmt.Printf("Command failed with: %s\n", err)
		_os.Exit(1)
	}
}
