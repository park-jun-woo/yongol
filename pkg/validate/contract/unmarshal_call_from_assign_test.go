//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestUnmarshalCallFromAssign — AssignStmt RHS 가 Unmarshal 호출이면 반환 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestUnmarshalCallFromAssign(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantNil bool
	}{
		{"json unmarshal", "err := json.Unmarshal(b, &v)", false},
		{"yaml unmarshal", "err := yaml.Unmarshal(b, &v)", false},
		{"non unmarshal", "err := json.Marshal(v)", true},
		{"not a call", "x := 1", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			got := unmarshalCallFromAssign(as)
			if (got == nil) != tt.wantNil {
				t.Fatalf("unmarshalCallFromAssign(%q) nil=%v, want %v", tt.body, got == nil, tt.wantNil)
			}
		})
	}
	if unmarshalCallFromAssign(nil) != nil {
		t.Fatal("nil assign should return nil")
	}
}
