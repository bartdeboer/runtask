package runtask

import (
	"embed"
	"fmt"
	"go/token"
	"io/fs"
	"os"
	"strings"
)

//go:embed embed_base.go
var embedFS embed.FS

func readTaskfile(files []string) ([]byte, error) {
	for _, file := range files {
		taskSrc, err := os.ReadFile(file)
		if err == nil {
			return taskSrc, nil
		}
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error reading Taskfile: %v", err)
		}
	}
	return nil, fmt.Errorf("could not find Taskfile")
}

func RunTask() error {

	fset := token.NewFileSet()

	if len(os.Args) < 2 || (len(os.Args) == 2 && os.Args[1] == "help") {
		basicHelp()
	}

	taskSrc, err := readTaskfile([]string{"Taskfile", "Taskfile.go", "tasks/Taskfile.go"})
	if err != nil {
		return err
	}

	ast2, err := parseSource(fset, string(taskSrc), "main")
	if err != nil {
		return fmt.Errorf("error parsing Taskfile: %v", err)
	}

	baseSrc, err := fs.ReadFile(embedFS, "embed_base.go")
	if err != nil {
		return fmt.Errorf("error loading base code: %v", err)
	}

	ast1, err := parseSource(fset, string(baseSrc), "main")
	if err != nil {
		return fmt.Errorf("error parsing base code: %v", err)
	}

	ast := mergeASTs(ast1, ast2)

	if len(os.Args) < 2 || (len(os.Args) == 2 && os.Args[1] == "help") {
		tasks, comments, _ := extractTasks(ast)
		tasksHelp(tasks, comments)
		return nil
	}

	taskName := strings.ToLower(os.Args[1])
	args := os.Args[2:]
	tasks, comments, taskArgs := extractTasks(ast)

	if taskName == "help" && len(os.Args) >= 3 {
		if _, ok := tasks[os.Args[2]]; !ok {
			return fmt.Errorf("no such task: %s", os.Args[2])
		}
		taskHelp(os.Args[2], comments, taskArgs)
		return nil
	}

	functionName, ok := tasks[taskName]
	if !ok {
		return fmt.Errorf("no such task: %s", taskName)
	}

	err = readDotEnv(".env")
	if err != nil {
		return fmt.Errorf("failed reading .env: %v", err)
	}

	// fmt.Println(astToString(fset, ast))

	i := newInterpreter()

	_, err = i.Eval(astToString(fset, ast))
	if err != nil {
		return fmt.Errorf("failed to evaluate taskfile: %v", err)
	}

	v, err := i.Eval(functionName)
	if err != nil {
		return fmt.Errorf("failed to reference task function %s: %v", functionName, err)
	}

	_, err = CallFunc(v, args)
	if err != nil {
		return fmt.Errorf("failed to call task function %s: %v", functionName, err)
	}
	return nil
}
