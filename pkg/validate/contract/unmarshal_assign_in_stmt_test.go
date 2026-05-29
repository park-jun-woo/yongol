//ff:func feature=validate-contract type=test control=selection topic=preserve-safety
//ff:what TestUnmarshalAssignInStmt — stmt 종류별 Unmarshal 호출 분류 디스패치 검증

package contract

import "testing"

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
			stmt := mustFirstStmt(t, tt.body)
			_, call, errName, kind := unmarshalAssignInStmt(stmt)
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
