//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestUnmarshalInAssignStmt — AssignStmt Unmarshal 호출의 kind/errName 분류 검증
package contract

import (
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
			assertUnmarshalInAssignStmt(t, tt.body, tt.wantNil, tt.wantKind, tt.wantErr)
		})
	}
}
