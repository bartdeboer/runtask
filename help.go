package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func basicHelp() {
	programName := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
	fmt.Println()
	fmt.Printf("Usage:  %s TASK [ARG...]\n", programName)
	fmt.Println()
}

func tasksHelp(functions map[string]string, comments map[string]string) {
	colWidth := 15
	totalWidth := 80
	fmt.Println("Available tasks:")
	fmt.Println()

	taskNames := make([]string, 0, len(functions))
	for taskName := range functions {
		taskNames = append(taskNames, taskName)
		if len(taskName) > colWidth {
			colWidth = len(taskName)
		}
	}
	sort.Strings(taskNames)

	formatStr := fmt.Sprintf("   %%-%ds    %%s\n", colWidth)
	for _, taskName := range taskNames {
		comment := comments[taskName]
		if len(comment) > totalWidth-colWidth-4 {
			comment = comment[:totalWidth-colWidth-4-3] + "..."
		}
		fmt.Printf(formatStr, taskName, comment)
	}

	fmt.Println()
	fmt.Printf(formatStr, "help [task]", "Show details about a task")
	fmt.Println()
}

func taskHelp(taskName string, comments map[string]string, args map[string][]string) {

	programName := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")

	if taskName == "help" {
		fmt.Println()
		fmt.Printf("Usage:  %s help COMMAND\n", programName)
		fmt.Println()
		fmt.Println("Show details about a task")
		fmt.Println()
		return
	}

	var taskArgs []string

	comment, _ := comments[taskName]
	taskArgs, ok := args[taskName]

	if !ok {
		taskArgs = []string{}
	}

	fmt.Println()
	fmt.Printf("Usage:  %s %s %s\n", programName, taskName, strings.Join(taskArgs, " "))
	fmt.Println()
	fmt.Println(comment)
	fmt.Println()
}
