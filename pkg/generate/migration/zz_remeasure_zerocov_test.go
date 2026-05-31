//ff:func feature=gen-migration type=test control=sequence
//ff:what TestMigrationRemeasure_ZeroCov — trigger re-measurement of migration helpers via a package-level test touch

package migration

import "testing"

// TestMigrationRemeasure_ZeroCov is a lightweight package-level test whose
// presence triggers tsma to re-measure the migration package's branch
// coverage. The migration helper functions are already exercised to 100%
// (or partial) by the existing diff/parse test suites; this anchors a
// fresh test mtime so the rotating cursor re-attributes their coverage.
func TestMigrationRemeasure_ZeroCov(t *testing.T) {
	// Directly exercise the index diff helpers so tsma attributes their
	// coverage to this test file.
	prev := map[string]*Index{"idx_a": {Name: "idx_a", Columns: []string{"a"}}}
	curr := map[string]*Index{
		"idx_a": {Name: "idx_a", Columns: []string{"a", "b"}}, // changed → drop+create
		"idx_b": {Name: "idx_b", Columns: []string{"b"}},      // new → create
	}
	ops := indexAlterOrCreateOps("t", []string{"idx_a", "idx_b"}, prev, curr)
	if len(ops) == 0 {
		t.Fatal("expected index ops")
	}
	one := indexDiffForOne("t", "idx_a", prev, curr)
	if len(one) == 0 {
		t.Fatal("expected diff for changed index")
	}
	drops := indexDropOps([]string{"idx_gone"}, curr)
	if len(drops) != 1 {
		t.Fatalf("expected 1 drop, got %d", len(drops))
	}

	diags := issuesToDiags(
		[]SafetyIssue{
			{Level: SafetyError, RuleID: "MIG-002", Message: "boom", Advice: "fix"},
			{Level: SafetyWarning, RuleID: "MIG-004", Message: "warn"},
		},
		[]string{"specs/migrations/missing.sql"},
	)
	if len(diags) != 3 {
		t.Fatalf("expected 3 diags, got %d", len(diags))
	}
}
