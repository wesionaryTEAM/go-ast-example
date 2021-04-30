package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "traversal/parse_this.txt", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Imports:")
	for _, i := range node.Imports {
		fmt.Println(i.Path.Value)
	}

	fmt.Println("Comments:")
	for _, c := range node.Comments {
		fmt.Println(c.Text())
	}

	fmt.Println("Functions:")
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println("func name: ", fn.Name)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Println("return statement found on line.", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
			return true
		}
		return true
	})
}
