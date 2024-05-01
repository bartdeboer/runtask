package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"unicode"
)

func overridePackageDecl(code, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(code))
	var builder strings.Builder
	builder.WriteString("package " + name + "\n")

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "package ") {
			continue
		}

		builder.WriteString(line + "\n")
	}

	return builder.String()
}

func parseSource(fset *token.FileSet, src, name string) (*ast.File, error) {
	src = overridePackageDecl(src, name)
	return parser.ParseFile(fset, "", src, parser.ParseComments)
}

func astToString(fset *token.FileSet, a *ast.File) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, a); err != nil {
		fmt.Printf("printer.Fprint error: %v\n", err)
		return ""
	}

	// Reformat the code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("format.Source error: %v\n", err)
		fmt.Printf("%s\n", string(buf.Bytes()))
		return string(buf.Bytes())
	}

	return string(formattedCode)
}

func mergeASTs(files ...*ast.File) *ast.File {
	newImportDecl := &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{},
	}
	newFile := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			newImportDecl,
		},
	}
	for _, file := range files {
		for _, decl := range file.Decls {
			if importDecl, ok := decl.(*ast.GenDecl); ok && importDecl.Tok == token.IMPORT {
				for _, spec := range importDecl.Specs {
					if importSpec, ok := spec.(*ast.ImportSpec); ok {
						newImportDecl.Specs = append(newImportDecl.Specs, importSpec)
					}
				}
			} else {
				newFile.Decls = append(newFile.Decls, decl)
			}
		}
	}
	return newFile
}

func extractTasks(file *ast.File) (map[string]string, map[string]string, map[string][]string) {
	functions := make(map[string]string)
	comments := make(map[string]string)
	argNames := make(map[string][]string)

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {

			funcName := fn.Name.Name

			// Only handle "exported" functions
			if fn.Recv != nil || !unicode.IsUpper(rune(funcName[0])) {
				continue
			}
			taskName := strings.ToLower(funcName)
			functions[taskName] = funcName

			// Comments
			var comment string
			if fn.Doc != nil {
				comment = strings.TrimSpace(fn.Doc.Text())
			}
			comments[taskName] = comment

			// Argument
			var args []string
			if fn.Type.Params != nil {
				for _, param := range fn.Type.Params.List {
					for _, ident := range param.Names {
						args = append(args, ident.Name)
					}
				}
			}
			argNames[taskName] = args
		}
	}

	return functions, comments, argNames
}
