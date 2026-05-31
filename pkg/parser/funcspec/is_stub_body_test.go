//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"go/token"
	"testing"
)

func TestIsStubBody(t *testing.T) {
	fset := token.NewFileSet()
	tests := []struct {
		name string
		body string
		want bool
	}{
		{"empty", ``, true},
		{"panic", `panic("TODO")`, true},
		{"zero return", `return Resp{}, nil`, true},
		{"meaningful return", `return Resp{Status: 1}, nil`, false},
		{"multi stmt", "x := 1\nreturn Resp{}, nil", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStubBody(fset, parseBody(t, tt.body)); got != tt.want {
				t.Errorf("isStubBody(%q) = %v, want %v", tt.body, got, tt.want)
			}
		})
	}
}
