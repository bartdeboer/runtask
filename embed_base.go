package main

// Prevent conflicts
import (
	_fmt "fmt"
	_os "os"
	_exec "os/exec"
)

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

func env(vars map[string]string) {
	for key, val := range vars {
		if err := _os.Setenv(key, val); err != nil {
			_fmt.Printf("Error setting environment variableCommand failed with: %s\n", err)
		}
	}
}
