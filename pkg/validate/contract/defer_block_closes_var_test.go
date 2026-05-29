//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestDeferBlockClosesVar — defer 클로저 body 내 Close 호출 존재 여부 검증

package contract

import "testing"

func TestDeferBlockClosesVar(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		varName string
		want    bool
	}{
		{"direct close in block", "f.Close()", "f", true},
		{"close inside if", "if f != nil { f.Close() }", "f", true},
		{"no close", "f.Read(b)", "f", false},
		{"wrong var", "f.Close()", "g", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts := mustStmts(t, tt.body)
			if got := deferBlockClosesVar(stmts, tt.varName); got != tt.want {
				t.Fatalf("deferBlockClosesVar(%q, %q) = %v, want %v", tt.body, tt.varName, got, tt.want)
			}
		})
	}
}
