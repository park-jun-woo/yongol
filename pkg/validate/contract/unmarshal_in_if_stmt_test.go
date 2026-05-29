//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestUnmarshalInIfStmt — if-init Unmarshal 은 Discarded(가드됨) 로 분류

package contract

import (
	"go/ast"
	"testing"
)

func TestUnmarshalInIfStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind unmarshalKind
	}{
		{"if init unmarshal", "if err := json.Unmarshal(b, &v); err != nil { return }", false, unmarshalKindDiscarded},
		{"if no init", "if x > 0 { return }", true, unmarshalKindUnknown},
		{"if init non unmarshal", "if err := json.Marshal(v); err != nil { return }", true, unmarshalKindUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifs := mustFirstStmt(t, tt.body).(*ast.IfStmt)
			_, call, _, kind := unmarshalInIfStmt(ifs)
			if (call == nil) != tt.wantNil {
				t.Fatalf("call nil=%v, want %v", call == nil, tt.wantNil)
			}
			if kind != tt.wantKind {
				t.Errorf("kind = %v, want %v", kind, tt.wantKind)
			}
		})
	}
}
