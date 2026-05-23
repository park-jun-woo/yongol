//ff:func feature=validate type=test control=sequence topic=migration-hints
//ff:what Run — 전입력 nil/empty 시 빈 결과 + 복합 진단 집계 검증

package migration

import (
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestRun(t *testing.T) {
	t.Run("all nil/empty inputs return empty", func(t *testing.T) {
		tmp := t.TempDir()
		prev := &gmig.Schema{Tables: map[string]*gmig.Table{}}
		curr := &gmig.Schema{Tables: map[string]*gmig.Table{}}
		diags := Run(tmp, prev, curr, nil, nil, nil)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("aggregates diagnostics from multiple rules", func(t *testing.T) {
		tmp := t.TempDir()
		prev := &gmig.Schema{Tables: map[string]*gmig.Table{}}
		curr := &gmig.Schema{Tables: map[string]*gmig.Table{}}
		issues := []gmig.SafetyIssue{
			{RuleID: "MIG-002", Message: "not null"},
			{RuleID: "MIG-004", Message: "destructive"},
		}
		missing := []string{"db/migrate/001.sql"}

		diags := Run(tmp, prev, curr, nil, issues, missing)
		// Expect: 1 MIG-002 + 1 MIG-003 + 1 MIG-004 = 3
		if len(diags) != 3 {
			t.Fatalf("expected 3 diagnostics, got %d: %+v", len(diags), diags)
		}
	})
}
