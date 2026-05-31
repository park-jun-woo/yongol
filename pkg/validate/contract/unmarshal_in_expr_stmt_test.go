//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestUnmarshalInExprStmt — bare Unmarshal ExprStmt 는 Assigned(errName="") 로 분류
package contract

import (
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
			assertUnmarshalInExprStmt(t, tt.body, tt.wantNil, tt.wantKind)
		})
	}
}
