//ff:func feature=gen-splitter type=test control=sequence
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"testing"
)

func TestFuncTypeFor(t *testing.T) {
	if got := funcTypeFor(ToolOAPICodegen); got != "handler" {
		t.Errorf("oapi = %q, want handler", got)
	}
	if got := funcTypeFor(ToolSQLC); got != "query" {
		t.Errorf("sqlc = %q, want query", got)
	}
	if got := funcTypeFor(Tool("x")); got != "util" {
		t.Errorf("default = %q, want util", got)
	}
}
