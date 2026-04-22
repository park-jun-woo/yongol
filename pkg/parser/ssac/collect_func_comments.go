//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what 함수 위치 이전의 주석을 수집
package ssac

import (
	"go/ast"
	"go/token"
)

// collectFuncComments는 함수 위치 이전의 주석을 수집한다.
func collectFuncComments(f *ast.File, fnPos token.Pos) []*ast.Comment {
	var comments []*ast.Comment
	for _, cg := range f.Comments {
		if cg.End() < fnPos {
			comments = append(comments, cg.List...)
		}
	}
	return comments
}
