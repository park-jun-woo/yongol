//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestUnmarshalInAssignStmt — AssignStmt Unmarshal 호출의 kind/errName 분류 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestUnmarshalInAssignStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind unmarshalKind
		wantErr  string
	}{
		{"assigned err", "err := json.Unmarshal(b, &v)", false, unmarshalKindAssigned, "err"},
		{"blank discard", "_ = json.Unmarshal(b, &v)", false, unmarshalKindDiscarded, ""},
		{"not unmarshal", "err := json.Marshal(v)", true, unmarshalKindUnknown, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			_, call, errName, kind := unmarshalInAssignStmt(as)
			if (call == nil) != tt.wantNil {
				t.Fatalf("call nil=%v, want %v", call == nil, tt.wantNil)
			}
			if kind != tt.wantKind {
				t.Errorf("kind = %v, want %v", kind, tt.wantKind)
			}
			if errName != tt.wantErr {
				t.Errorf("errName = %q, want %q", errName, tt.wantErr)
			}
		})
	}
}
