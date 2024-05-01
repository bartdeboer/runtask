package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

type TestTask struct {
	Taskfile string
	OsArgs   []string
	ChDir    string
	Contains string
	Expected string
	Error    string
}

func runTestTask(task TestTask, t *testing.T) {
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %s", err)
	}
	err = os.Chdir(task.ChDir)
	if err != nil {
		t.Fatalf("Failed to change working directory: %s", err)
	}
	oldOsArgs := os.Args
	os.Args = task.OsArgs

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %s", err)
	}
	oldStdout := os.Stdout
	os.Stdout = w

	defer func() {
		w.Close()
		os.Stdout = oldStdout
		os.Args = oldOsArgs
		err = os.Chdir(oldWd)
		if err != nil {
			t.Errorf("Failed to restore working directory: %s", err)
		}
	}()

	// Execute the task
	taskErr := runTask()

	if err != nil {
		t.Fatalf("Failed testing task: %s", err)
	}

	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %s", err)
	}

	if task.Error != "" && fmt.Sprintf("%s", taskErr) != task.Error {
		t.Errorf("Expected error '%s', got '%s'", task.Error, taskErr)
	}

	if task.Expected != "" && string(out) != task.Expected {
		t.Errorf("Expected '%s', got '%s'", task.Expected, string(out))
	}

	if task.Contains != "" && strings.Contains(string(out), task.Contains) {
		t.Errorf("Expected to contain '%s', got '%s'", task.Contains, string(out))
	}
}

func TestTaskRunner(t *testing.T) {

	// t1 notexist
	runTestTask(TestTask{
		ChDir:  "tests/t1/",
		OsArgs: []string{"runtask", "notexist"},
		Error:  "no such task: notexist",
	}, t)

	// t1
	runTestTask(TestTask{
		ChDir:  "tests/",
		OsArgs: []string{"runtask"},
		Error:  "could not find Taskfile",
	}, t)

	// t1 simple
	runTestTask(TestTask{
		ChDir:    "tests/t1/",
		OsArgs:   []string{"runtask", "simple"},
		Expected: "Hello, World!\n",
	}, t)

	// t1 simple
	runTestTask(TestTask{
		ChDir:  "tests/t1/",
		OsArgs: []string{"runtask", "simple", "arg1"},
		Error:  "failed to call task function Simple: wrong parameter count",
	}, t)

	// t1 simplearg
	runTestTask(TestTask{
		ChDir:  "tests/t1/",
		OsArgs: []string{"runtask", "simplearg"},
		Error:  "failed to call task function SimpleArg: wrong parameter count",
	}, t)

	// t1 simplearg
	runTestTask(TestTask{
		ChDir:    "tests/t1/",
		OsArgs:   []string{"runtask", "simplearg", "John"},
		Expected: "Hello, John!\n",
	}, t)

	// t1 simplearg
	runTestTask(TestTask{
		ChDir:  "tests/t1/",
		OsArgs: []string{"runtask", "simplearg", "John", "Jane"},
		Error:  "failed to call task function SimpleArg: wrong parameter count",
	}, t)

	// t1 variadic
	runTestTask(TestTask{
		ChDir:  "tests/t1/",
		OsArgs: []string{"runtask", "variadic"},
		Error:  "failed to call task function Variadic: wrong parameter count",
	}, t)

	// t1 variadic
	runTestTask(TestTask{
		ChDir:    "tests/t1/",
		OsArgs:   []string{"runtask", "variadic", "Jane"},
		Expected: "Hello, Jane []!\n",
	}, t)

	// t1 variadic
	runTestTask(TestTask{
		ChDir:    "tests/t1/",
		OsArgs:   []string{"runtask", "variadic", "John", "Jane", "Johnie"},
		Expected: "Hello, John [Jane, Johnie]!\n",
	}, t)

	// t1 add1
	runTestTask(TestTask{
		ChDir:    "tests/t1/",
		OsArgs:   []string{"runtask", "add1", "1.567", "1", "2", "3", "4", "5"},
		Expected: "1.567 15\n",
	}, t)

	// t1 add2
	runTestTask(TestTask{
		ChDir:    "tests/t1/",
		OsArgs:   []string{"runtask", "add2", "1.567", "1", "2", "3", "4", "5"},
		Expected: "1.567 15\n",
	}, t)

	// t1 add2
	runTestTask(TestTask{
		ChDir:  "tests/t1/",
		OsArgs: []string{"runtask", "add2", "1.567", "1", "2", "3", "string", "5"},
		Error:  `failed to call task function Add2: converting argument 4: converting string to int: strconv.ParseInt: parsing "string": invalid syntax`,
	}, t)

	// t2/simple
	runTestTask(TestTask{
		ChDir:    "tests/t2/",
		OsArgs:   []string{"runtask", "simple"},
		Expected: "FOO: Hello, Jane!\nBAR: Hello, John!\n",
	}, t)

}
