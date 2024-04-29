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
)

func ensurePackageDeclaration(code, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(code))
	hasPackage := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "package ") {
			hasPackage = true
			break
		}
		break
	}

	if hasPackage {
		return code // Return original code if package declaration exists
	} else {
		return "package " + name + "\n\n" + code // Prepend 'package main' if no package declaration
	}
}

func parseSource(fset *token.FileSet, src string) (*ast.File, error) {
	src = ensurePackageDeclaration(src, "main")
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

func findTasks(file *ast.File) ([]string, map[string]string, map[string][]string) {
	var functions []string
	comments := make(map[string]string)
	argNames := make(map[string][]string) // Added to store argument names

	// Loop through the declarations in the file
	for _, decl := range file.Decls {
		// Check if the declaration is a function
		if fn, ok := decl.(*ast.FuncDecl); ok {
			// Skip methods and functions named "r"
			if fn.Recv != nil || fn.Name.Name == "run" || fn.Name.Name == "env" {
				continue
			}
			name := fn.Name.Name
			functions = append(functions, name)

			// Process documentation comments
			var comment string
			if fn.Doc != nil {
				for _, astComment := range fn.Doc.List {
					comment += " " + strings.TrimSpace(astComment.Text)
				}
			}
			comments[name] = comment

			// Extracting argument names
			var args []string
			if fn.Type.Params != nil {
				for _, param := range fn.Type.Params.List {
					for _, ident := range param.Names {
						args = append(args, ident.Name)
					}
				}
			}
			argNames[name] = args
		}
	}

	return functions, comments, argNames
}
