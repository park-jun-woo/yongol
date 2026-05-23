//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-safety
//ff:what emitByRule — 빈 입력/매칭/비매칭 RuleID 필터 + 레벨·메시지 검증

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestEmitByRule(t *testing.T) {
	issues := []gmig.SafetyIssue{
		{RuleID: "MIG-002", Message: "drop column detected", Advice: "add safety hint"},
		{RuleID: "MIG-004", Message: "rename table detected", Advice: "use alias"},
		{RuleID: "MIG-002", Message: "drop table detected", Advice: "backup first"},
	}

	tests := []struct {
		name      string
		issues    []gmig.SafetyIssue
		ruleID    string
		lvl       diagnostic.Level
		wantCount int
		wantSub   string
	}{
		{
			name:      "nil issues returns empty",
			issues:    nil,
			ruleID:    "MIG-002",
			lvl:       diagnostic.LevelError,
			wantCount: 0,
		},
		{
			name:      "no matching rule returns empty",
			issues:    issues,
			ruleID:    "MIG-005",
			lvl:       diagnostic.LevelError,
			wantCount: 0,
		},
		{
			name:      "MIG-002 matches two issues",
			issues:    issues,
			ruleID:    "MIG-002",
			lvl:       diagnostic.LevelError,
			wantCount: 2,
			wantSub:   "MIG-002",
		},
		{
			name:      "MIG-004 matches one issue",
			issues:    issues,
			ruleID:    "MIG-004",
			lvl:       diagnostic.LevelWarning,
			wantCount: 1,
			wantSub:   "MIG-004",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := emitByRule(tt.issues, tt.ruleID, tt.lvl)
			assertEmitByRule(t, diags, tt.wantCount, tt.wantSub, tt.lvl)
		})
	}
}
