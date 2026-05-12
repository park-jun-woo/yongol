//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TestResolveInputType_CurrentUserField — currentUser 필드 접근의 claim 타입 해석 검증

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// TestResolveInputType_CurrentUserField verifies that resolveInputType returns
// the registered claim type for currentUser.<Field> expressions.
func TestResolveInputType_CurrentUserField(t *testing.T) {
	g := &rule.Ground{
		Types: map[string]string{
			"Manifest.claim.OrgID":  "pgtype.UUID",
			"Manifest.claim.UserID": "int64",
		},
	}
	tests := []struct {
		value string
		want  string
	}{
		{"currentUser.OrgID", "pgtype.UUID"},
		{"currentUser.UserID", "int64"},
		{"currentUser.Unknown", ""},
		{"org.Name", ""},
	}
	for _, tt := range tests {
		got := resolveInputType(g, "anyFunc", tt.value)
		if got != tt.want {
			t.Errorf("resolveInputType(g, %q, %q) = %q, want %q", "anyFunc", tt.value, got, tt.want)
		}
	}
}
