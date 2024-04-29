package main

import (
	"fmt"
	"strings"
)

func displayTasks2(functions []string, comments map[string]string, argNames map[string][]string) {
	for _, functionName := range functions {
		fmt.Printf("Task: %s\n", functionName)
		if args, ok := argNames[functionName]; ok && len(args) > 0 {
			fmt.Printf("Arguments: %s\n", strings.Join(args, ", "))
		} else {
			fmt.Printf("Arguments: None\n")
		}
		if comment, ok := comments[functionName]; ok && len(strings.TrimSpace(comment)) > 0 {
			fmt.Printf("Description:%s\n", comment)
		} else {
			fmt.Printf("Description: No description available.\n")
		}
		fmt.Println(strings.Repeat("-", 80)) // Print a separator line
	}
}

func displayTasks(functions []string, comments map[string]string, argNames map[string][]string) {
	// Determine the maximum width needed for the command and its arguments
	maxWidth := 20
	colWidth := 15
	for _, functionName := range functions {
		command := functionName
		if args, ok := argNames[functionName]; ok {
			command += " " + strings.Join(args, " ")
		}
		if len(command) > colWidth {
			colWidth = len(command)
		}
	}

	// Format and display each command and its description
	for _, functionName := range functions {
		command := functionName
		if args, ok := argNames[functionName]; ok {
			command += " " + strings.Join(args, " ")
		}
		comment, _ := comments[functionName]
		comment = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comments[functionName]), "//"))

		// Print command with padding to align comments
		if colWidth > maxWidth {
			formatStr := fmt.Sprintf("   %%-%ds    %%s\n", maxWidth)
			fmt.Printf("   %s\n", command)
			fmt.Printf(formatStr, "", comment)
			fmt.Println()
		} else {
			formatStr := fmt.Sprintf("   %%-%ds    %%s\n", colWidth)
			fmt.Printf(formatStr, command, comment)
		}
	}
}
