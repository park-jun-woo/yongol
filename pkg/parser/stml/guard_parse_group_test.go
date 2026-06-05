//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardGroup — "(" guard ")" 괄호 그룹 파싱 (정상/내부 에러/닫힘 누락) 검증

package stml

import "testing"

func TestParseGuardGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantPath string
	}{
		{name: "valid group", input: "(a.b=x)", wantPath: "a.b"},
		{name: "inner error", input: "(&&)", wantErr: true},
		{name: "missing close paren", input: "(a.b=x", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardGroup(t, tt.input, tt.wantErr, tt.wantPath)
		})
	}
}
