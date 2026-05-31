//ff:func feature=ssac-parse type=test-helper control=iteration dimension=1
//ff:what exerciseFileComments — 파일의 각 comment group 에 parseComments 직접 호출 헬퍼
package ssac

import (
	"go/ast"
	"go/token"
)

// exerciseFileComments calls parseComments on each comment group of the file.
func exerciseFileComments(fset *token.FileSet, f *ast.File) {
	for _, cg := range f.Comments {
		parseComments(fset, cg.List)
	}
}
