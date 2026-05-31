//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractJSONTag(t *testing.T) {
	field := func(tag string) *ast.Field {
		var f ast.Field
		if tag != "" {
			f.Tag = &ast.BasicLit{Kind: token.STRING, Value: "`" + tag + "`"}
		}
		return &f
	}
	if got := extractJSONTag(field(`json:"user_id"`)); got != "user_id" {
		t.Errorf("got %q, want user_id", got)
	}
	if got := extractJSONTag(field(`json:"name,omitempty"`)); got != "name" {
		t.Errorf("got %q, want name (option stripped)", got)
	}
	if got := extractJSONTag(field(`json:"-"`)); got != "" {
		t.Errorf("json:- → %q, want empty", got)
	}
	if got := extractJSONTag(field("")); got != "" {
		t.Errorf("no tag → %q, want empty", got)
	}
	if got := extractJSONTag(field(`xml:"x"`)); got != "" {
		t.Errorf("no json tag → %q, want empty", got)
	}
}
