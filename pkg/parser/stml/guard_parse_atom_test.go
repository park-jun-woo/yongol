//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardAtom — atom (group / lifecycle / compare / 에러) 파싱 분기 검증

package stml

import "testing"

func TestParseGuardAtom(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantKind GuardKind
		wantPath string
	}{
		{name: "group", input: "(a.b=x)", wantKind: GuardGroup, wantPath: "a.b"},
		{name: "lifecycle", input: "a.b.loading", wantKind: GuardLifecycle, wantPath: "a.b"},
		{name: "compare", input: "a.b=x", wantKind: GuardCompare, wantPath: "a.b"},
		{name: "missing operator", input: "a.b", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardAtom(t, tt.input, tt.wantErr, tt.wantKind, tt.wantPath)
		})
	}
}
