package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"go/ast"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "simple",
	Doc: "文字を解析します",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer}, // inspect.Analyzerを利用
}

func run(pass *analysis.Pass) (interface{}, error) {
	// inspect.Analyzerの結果を取得
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// ASTのすべての識別子(ident)を探索
	inspect.Preorder([]ast.Node{&ast.Ident{}}, func(n ast.Node) {
		if ident, ok := n.(*ast.Ident); ok {
			if ident.Name == "gopher" {
				pass.Reportf(ident.Pos(), "gopherという文字が検出されました")
			}
		}
	})
	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer) 
}