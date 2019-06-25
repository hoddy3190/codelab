package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()

	// ファイル単位で構文解析を行う
	f, err := parser.ParseFile(fset, "gopher.go", nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// 抽象構文木を深さ優先で探索する
	ast.Inspect(f, func(n ast.Node) bool {

		// 関数呼び出し
		ident, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		// Selector
		ident2, ok := ident.Fun.(*ast.SelectorExpr)
		if !ok {
		    return true
		}

		fmt.Println(fset.Position(ident2.Pos()))

		return true
	})
}
