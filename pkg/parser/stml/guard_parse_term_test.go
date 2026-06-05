//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardTerm — term := "!"? atom 파싱 (부정/비부정/부정 후 에러) 검증

package stml

import "testing"

func TestParseGuardTerm(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantKind GuardKind
	}{
		{name: "negation", input: "!a.b=x", wantKind: GuardUnary},
		{name: "plain atom", input: "a.b=x", wantKind: GuardCompare},
		{name: "negation with error", input: "!&&", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardTerm(t, tt.input, tt.wantErr, tt.wantKind)
		})
	}
}
