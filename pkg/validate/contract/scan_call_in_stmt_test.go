//ff:func feature=validate-contract type=test control=selection topic=preserve-safety
//ff:what TestScanCallInStmt — stmt 종류별 Scan 호출 분류 디스패치 검증

package contract

import "testing"

func TestScanCallInStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind scanKind
		wantErr  string
	}{
		{"if init", "if err := row.Scan(&x); err != nil { return }", false, scanKindDiscarded, ""},
		{"assign", "err := row.Scan(&x)", false, scanKindAssigned, "err"},
		{"blank assign", "_ = row.Scan(&x)", false, scanKindDiscarded, ""},
		{"bare expr", "row.Scan(&x)", false, scanKindAssigned, ""},
		{"unrelated", "x := 1", true, scanKindUnknown, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := mustFirstStmt(t, tt.body)
			call, errName, kind := scanCallInStmt(stmt)
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
