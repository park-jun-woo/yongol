//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what collectShapesFromFile — 파싱된 파일의 TYPE 선언을 훑어 JSONResponse 래퍼를 shapes 맵에 분류 적재

package ssac

import (
	"go/ast"
	"go/token"
	"strings"
)

// collectShapesFromFile walks the TYPE declarations of one parsed file and
// records every `type <Name>JSONResponse ...` whose shape classifyTypeSpec
// recognises into shapes (keyed by the wrapper type name). Unrecognised
// shapes are left out so the alias fallback applies at emit time. Splitting
// this out of classifyResponseShapes keeps that function's loop nesting flat.
func collectShapesFromFile(file *ast.File, shapes map[string]respShape) {
	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || !strings.HasSuffix(ts.Name.Name, "JSONResponse") {
				continue
			}
			if shape, ok := classifyTypeSpec(ts); ok {
				shapes[ts.Name.Name] = shape
			}
		}
	}
}
