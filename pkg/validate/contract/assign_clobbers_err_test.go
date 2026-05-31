//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestAssignClobbersErr — stmt 가 errName 을 새 값으로 덮어쓰는지 판정 검증
package contract

import (
	"testing"
)

func TestAssignClobbersErr(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		errName string
		want    bool
	}{
		{"clobber default", "err = other()", "", true},
		{"err nil not clobber", "err = nil", "", false},
		{"untracked assign", "x = other()", "", false},
		{"named clobber", "dbErr = other()", "dbErr", true},
		{"non assign stmt", "return", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := mustFirstStmt(t, tt.body)
			if got := assignClobbersErr(stmt, tt.errName); got != tt.want {
				t.Fatalf("assignClobbersErr(%q, %q) = %v, want %v", tt.body, tt.errName, got, tt.want)
			}
		})
	}
}
