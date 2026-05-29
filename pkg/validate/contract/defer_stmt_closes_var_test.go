//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestDeferStmtClosesVar — defer Close (직접 또는 closure) 인지 판정 검증

package contract

import "testing"

func TestDeferStmtClosesVar(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		varName string
		want    bool
	}{
		{"direct defer close", "defer f.Close()", "f", true},
		{"closure defer close", "defer func() { f.Close() }()", "f", true},
		{"closure guarded close", "defer func() { if f != nil { f.Close() } }()", "f", true},
		{"defer wrong var", "defer f.Close()", "g", false},
		{"defer non close", "defer f.Read(b)", "f", false},
		{"not a defer", "f.Close()", "f", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := mustFirstStmt(t, tt.body)
			if got := deferStmtClosesVar(stmt, tt.varName); got != tt.want {
				t.Fatalf("deferStmtClosesVar(%q, %q) = %v, want %v", tt.body, tt.varName, got, tt.want)
			}
		})
	}
}
