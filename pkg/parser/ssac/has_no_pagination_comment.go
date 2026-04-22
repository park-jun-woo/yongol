//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what hasNoPaginationComment — 함수 주석들 중 @no-pagination 지시문이 있는지 검사
package ssac

import (
	"go/ast"
	"strings"
)

// hasNoPaginationComment reports whether any line in the given comment slice
// matches the bare "@no-pagination" directive.
func hasNoPaginationComment(comments []*ast.Comment) bool {
	for _, c := range comments {
		line := strings.TrimPrefix(c.Text, "//")
		line = strings.TrimSpace(line)
		if line == "@no-pagination" {
			return true
		}
	}
	return false
}
