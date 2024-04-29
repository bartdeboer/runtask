package main

import (
	"embed"
	"fmt"
	"go/token"
	"io/fs"
	"os"
	"slices"
)

//go:embed embed_base.go
var embedFS embed.FS

func main() {
	fset := token.NewFileSet()

	baseSrc, err := fs.ReadFile(embedFS, "embed_base.go")
	if err != nil {
		fmt.Printf("Error loading base source: %v\n", err)
		os.Exit(1)
	}

	a1, err := parseSource(fset, string(baseSrc))
	if err != nil {
		fmt.Printf("Error parsing base code: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stat("taskfile")
	if os.IsNotExist(err) {
		fmt.Printf("Could not find taskfile: %v\n", err)
		os.Exit(1)
	}

	taskSrc, err := os.ReadFile("taskfile")
	if err != nil {
		fmt.Printf("Error reading taskfile: %v\n", err)
		os.Exit(1)
	}

	a2, err := parseSource(fset, string(taskSrc))
	if err != nil {
		fmt.Printf("Error parsing taskfile: %v\n", err)
		os.Exit(1)
	}

	a := mergeASTs(a1, a2)

	if len(os.Args) < 2 {
		fmt.Println("Available tasks:")
		displayTasks(findTasks(a))
		return
	}

	task := os.Args[1]
	args := os.Args[2:]

	tasks, _, _ := findTasks(a)

	if !slices.Contains(tasks, task) {
		fmt.Printf("No such task: %s\n", task)
		os.Exit(1)
	}

	i := newInterpreter()

	_, err = i.Eval(astToString(fset, a))
	if err != nil {
		fmt.Printf("Failed to evaluate taskfile: %v\n", err)
		os.Exit(1)
	}

	v, err := i.Eval(task)
	if err != nil {
		fmt.Printf("Failed to reference task function %s: %v\n", task, err)
		os.Exit(1)
	}

	_, err = callFunc(v, args)
	if err != nil {
		fmt.Printf("Failed to execute task: %v\n", err)
		os.Exit(1)
	}
}
