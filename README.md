### Go AST Example
Sample and notes for go abstract tree parsing and modification.

#### Reference Tutorials:
- Go Ast Traversal [here](https://zupzup.org/go-ast-traversal/)
- Go Ast Manipulation [here](https://www.zupzup.org/ast-manipulation-go/)

#### Packages
- `go/token`: Defines constants representing the lexical tokens of the go programming language and basic operations on tokens (printing, predicates)
  `"gopher" -> token.String`
- `go/scanner` Implements scanner for go source text. Takes `[]byte` and tokenize through repeated calls to scan method.
- `go/ast` declares types used to represent abstract syntax tree.
- `go/parser` Parses go source file and outputs AST. Parser package uses scanner and token package itself.
- `go/printer` prints ast in human readable format

`ast.Inspect` inspect traverses the ast node. 