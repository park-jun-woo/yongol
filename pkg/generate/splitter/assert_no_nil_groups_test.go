//ff:func feature=gen-splitter type=test-helper control=iteration dimension=1
//ff:what assertNoNilGroups — preserveComments 결과에 nil CommentGroup 이 없는지 검증 헬퍼
package splitter

import (
	"go/ast"
	"testing"
)

// assertNoNilGroups asserts no nil group leaked into the preserved comment list.
func assertNoNilGroups(t *testing.T, groups []*ast.CommentGroup) {
	t.Helper()
	for _, g := range groups {
		if g == nil {
			t.Fatal("nil group leaked into output")
		}
	}
}
