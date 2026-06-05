//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardExpr — term ((&&|||) term)* 좌결합 파싱 (단일/결합/좌·우 에러) 검증

package stml

import "testing"

func TestParseGuardExpr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantKind GuardKind
		wantOp   string
	}{
		{name: "single term", input: "a.b=x", wantKind: GuardCompare},
		{name: "and combination", input: "a.b=x && c.d=y", wantKind: GuardBinary, wantOp: "&&"},
		{name: "or combination", input: "a.b=x || c.d=y", wantKind: GuardBinary, wantOp: "||"},
		{name: "left error", input: "&& a.b=x", wantErr: true},
		{name: "right error", input: "a.b=x && &&", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardExpr(t, tt.input, tt.wantErr, tt.wantKind, tt.wantOp)
		})
	}
}
