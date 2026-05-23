//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-safety
//ff:what Mig002NotNullWithoutBackfill — MIG-002 필터 + ERROR 레벨 검증

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMig002NotNullWithoutBackfill(t *testing.T) {
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
			name: "no MIG-002 issues returns empty",
			issues: []gmig.SafetyIssue{
				{RuleID: "MIG-004", Message: "other issue"},
			},
			wantCount: 0,
		},
		{
			name: "MIG-002 issue emitted as ERROR",
			issues: []gmig.SafetyIssue{
				{RuleID: "MIG-002", Message: "NOT NULL without backfill"},
			},
			wantCount: 1,
		},
		{
			name: "multiple MIG-002 issues all emitted",
			issues: []gmig.SafetyIssue{
				{RuleID: "MIG-002", Message: "col1 not null"},
				{RuleID: "MIG-004", Message: "other"},
				{RuleID: "MIG-002", Message: "col2 not null"},
			},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := Mig002NotNullWithoutBackfill(tt.issues)
			assertDiagsLevel(t, diags, tt.wantCount, diagnostic.LevelError, "MIG-002")
		})
	}
}
