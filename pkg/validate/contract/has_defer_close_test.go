//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestHasDeferClose — 블록 내 defer varName.Close() 존재 여부 판정 검증

package contract

import "testing"

func TestHasDeferClose(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		varName string
		want    bool
	}{
		{"present", "x := 1\ndefer f.Close()\nreturn\n", "f", true},
		{"closure present", "defer func() { f.Close() }()\nreturn\n", "f", true},
		{"absent", "x := 1\nreturn\n", "f", false},
		{"wrong var", "defer f.Close()\n", "g", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts := mustStmts(t, tt.body)
			if got := hasDeferClose(stmts, tt.varName); got != tt.want {
				t.Fatalf("hasDeferClose(%q, %q) = %v, want %v", tt.body, tt.varName, got, tt.want)
			}
		})
	}
}
