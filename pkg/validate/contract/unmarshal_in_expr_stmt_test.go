//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestUnmarshalInExprStmt — bare Unmarshal ExprStmt 는 Assigned(errName="") 로 분류

package contract

import (
	"go/ast"
	"testing"
)

func TestUnmarshalInExprStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind unmarshalKind
	}{
		{"bare unmarshal", "json.Unmarshal(b, &v)", false, unmarshalKindAssigned},
		{"non unmarshal expr", "json.Marshal(v)", true, unmarshalKindUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := mustFirstStmt(t, tt.body).(*ast.ExprStmt)
			_, call, errName, kind := unmarshalInExprStmt(es)
			if (call == nil) != tt.wantNil {
				t.Fatalf("call nil=%v, want %v", call == nil, tt.wantNil)
			}
			if kind != tt.wantKind {
				t.Errorf("kind = %v, want %v", kind, tt.wantKind)
			}
			if errName != "" {
				t.Errorf("errName = %q, want empty", errName)
			}
		})
	}
}
