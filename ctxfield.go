package ctxfield

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"go/types"
)

// context.Contextをフィールドに持つ構造体を検出するAnalyzer
var Analyzer = &analysis.Analyzer{
	Name: "ctxfield",
	Doc: "context.Contextをフィールドに持つ構造体を検出します",
	Run: run,
}

// Run処理を実装
func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// 型宣言を探す
			if ts, ok := n.(*ast.TypeSpec); ok {
				// 構造体かどうかを確認する
				if st, ok := ts.Type.(*ast.StructType); ok {
					for _, field := range st.Fields.List {
						// フィールドの型情報を取得
						if fieldType, ok := pass.TypesInfo.Types[field.Type]; ok {
							// 型がcontext.Contextかどうかを確認する
							if named, ok := fieldType.Type.(*types.Named); ok {
								if named.Obj().Pkg() != nil && named.Obj().Pkg().Path() == "context" && named.Obj().Name() == "Context" {
									pass.Reportf(field.Pos(), "構造体のフィールドにcontext.Contextが含まれてしまっています: %s", field.Names[0].Name)
								}
							}
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}



