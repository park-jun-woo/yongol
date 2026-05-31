//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"go/ast"
	"testing"
)

func TestHasNoPaginationComment(t *testing.T) {
	yes := []*ast.Comment{
		{Text: "// some doc"},
		{Text: "// @no-pagination"},
	}
	if !hasNoPaginationComment(yes) {
		t.Error("expected @no-pagination detected")
	}
	no := []*ast.Comment{{Text: "// @something-else"}}
	if hasNoPaginationComment(no) {
		t.Error("should not detect @no-pagination")
	}
	if hasNoPaginationComment(nil) {
		t.Error("nil comments → false")
	}
}
