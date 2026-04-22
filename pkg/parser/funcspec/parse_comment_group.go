//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what 코멘트 그룹에서 @func, @error, @description 어노테이션을 추출한다
package funcspec

import (
	"go/ast"
	"go/token"
	"strings"
)

func parseCommentGroup(cg *ast.CommentGroup, fset *token.FileSet, spec *FuncSpec) {
	for _, c := range cg.List {
		line := strings.TrimPrefix(c.Text, "//")
		line = strings.TrimSpace(line)
		lineNum := fset.Position(c.Pos()).Line
		applyAnnotation(line, lineNum, spec)
	}
}
