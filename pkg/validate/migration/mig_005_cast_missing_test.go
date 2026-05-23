//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-safety
//ff:what Mig005CastMissing — MIG-005 필터 + WARNING 레벨 검증

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMig005CastMissing(t *testing.T) {
	tests := []struct {
		name      string
		issues    []gmig.SafetyIssue
		wantCount int
	}{
		{
			name:      "nil issues returns empty",
			issues:    nil,
			wantCount: 0,
		},
		{
			name: "no MIG-005 issues returns empty",
			issues: []gmig.SafetyIssue{
				{RuleID: "MIG-002", Message: "other"},
			},
			wantCount: 0,
		},
		{
			name: "MIG-005 issues emitted as WARNING",
			issues: []gmig.SafetyIssue{
				{RuleID: "MIG-005", Message: "type change without cast"},
			},
			wantCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := Mig005CastMissing(tt.issues)
			assertDiagsLevel(t, diags, tt.wantCount, diagnostic.LevelWarning, "MIG-005")
		})
	}
}
