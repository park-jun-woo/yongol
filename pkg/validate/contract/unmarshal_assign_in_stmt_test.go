//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestUnmarshalAssignInStmt — stmt 종류별 Unmarshal 호출 분류 디스패치 검증
package contract

import (
	"testing"
)

func TestUnmarshalAssignInStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind unmarshalKind
		wantErr  string
	}{
		{"if init", "if err := json.Unmarshal(b, &v); err != nil { return }", false, unmarshalKindDiscarded, ""},
		{"assign", "err := json.Unmarshal(b, &v)", false, unmarshalKindAssigned, "err"},
		{"blank assign", "_ = json.Unmarshal(b, &v)", false, unmarshalKindDiscarded, ""},
		{"bare expr", "json.Unmarshal(b, &v)", false, unmarshalKindAssigned, ""},
		{"unrelated", "x := 1", true, unmarshalKindUnknown, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertUnmarshalAssignInStmt(t, tt.body, tt.wantNil, tt.wantKind, tt.wantErr)
		})
	}
}
