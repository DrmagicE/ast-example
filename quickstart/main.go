package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

func main() {
	fset := token.NewFileSet()
	path, _ := filepath.Abs("./demo.go")
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	ast.Print(fset, f)
}
