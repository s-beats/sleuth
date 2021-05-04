package sleuth

import (
	"go/ast"
	"log"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "sleuth is a tool that can detect when you have `capacity` and you `append`"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "sleuth",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var slices map[string]int

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			call, ok := n.Rhs[0].(*ast.CallExpr)
			if ok {
				fun, ok := call.Fun.(*ast.Ident)
				if ok {
					switch fun.Name {
					case "make":
						args, ok := n.Rhs[0].(*ast.CallExpr)
						if ok {
							switch args.Args[0].(type) {
							case *ast.ArrayType:
								lenVal, ok := args.Args[1].(*ast.BasicLit)
								if ok {
									iLen, err := strconv.Atoi(lenVal.Value)
									if err != nil {
										log.Println(err)
										return
									}
									if iLen != 0 {
										lhs, ok := n.Lhs[0].(*ast.Ident)
										if ok {
											if slices == nil {
												slices = make(map[string]int)
											}
											slices[lhs.Name] = iLen
										}
									}
								}
							}
						}
					case "append":
						lhs, ok := n.Lhs[0].(*ast.Ident)
						if ok {
							if _, ok := slices[lhs.Name]; ok {
								pass.Reportf(n.Pos(), "sleuth detects illegal")
								delete(slices, lhs.Name)
								if _, ok := slices[lhs.Name]; ok {
									log.Printf("failed to delete key: %v, value: %v", lhs.Name, slices[lhs.Name])
								}
							}
						}
					}
				}
			}
		}
	})

	return nil, nil
}
