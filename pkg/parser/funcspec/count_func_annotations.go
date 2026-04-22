//ff:func feature=funcspec type=parser control=iteration dimension=2
//ff:what countFuncAnnotations — 파일 주석에서 @func 어노테이션 개수 집계
package funcspec

import (
	"go/ast"
	"strings"
)

// countFuncAnnotations counts "@func " occurrences across all file-level
// comment groups. Used to enforce "one @func per file" (BUG002).
func countFuncAnnotations(comments []*ast.CommentGroup) int {
	count := 0
	for _, cg := range comments {
		for _, c := range cg.List {
			line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			if strings.HasPrefix(line, "@func ") {
				count++
			}
		}
	}
	return count
}
