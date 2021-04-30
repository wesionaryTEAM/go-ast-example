package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	log.Println("Parse to ast")
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "manipulation/parse_this.txt", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	comments := []*ast.CommentGroup{}
	ast.Inspect(node, func(n ast.Node) bool {
		c, ok := n.(*ast.CommentGroup)
		if ok {
			comments = append(comments, c)
		}

		fn, ok := n.(*ast.FuncDecl)
		if ok {
			if fn.Name.IsExported() && fn.Doc.Text() == "" {
				log.Printf("export function declaration without documentation found on line %d: \n\t%s\n", fset.Position(fn.Pos()).Line, fn.Name.Name)

				comment := &ast.Comment{
					Text:  "// TODO: document exported function",
					Slash: fn.Pos() - 1,
				}

				cg := &ast.CommentGroup{
					List: []*ast.Comment{comment},
				}

				fn.Doc = cg
			}
		}

		return true
	})

	node.Comments = comments
	f, _ := os.Create("manipulation/new.txt")
	defer f.Close()
	if err := printer.Fprint(f, fset, node); err != nil {
		log.Fatal(err)
	}
}
