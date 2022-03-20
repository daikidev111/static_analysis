package error_handling

import (
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "error_handling suggests to use sync.ErrGroup instead of sync.WaitGroup as sync.ErrGroup"

var Analyzer = &analysis.Analyzer{
	Name: "error_handling",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "sample.go", nil, 0)
	ast.Print(fset, f)
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GoStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GoStmt:
			fn, ok := n.Call.Fun.(*ast.FuncLit)
			if !ok {
				return
			}
			ls := fn.Type.Params.List
			for _, l := range ls { // inside lists
				chTyp, ok := l.Type.(*ast.ChanType)
				if !ok { // if it fails to cast, ,the it will cause panic
					return
				}
				if chTyp.Arrow != 0 {
					return
				}
				val, ok := chTyp.Value.(*ast.Ident)
				if !ok {
					return
				}
				if val.Name == "error" {
					pass.Reportf(val.NamePos, "identifier is goroutine")
				}
			}

		}
	})

	return nil, nil
}
