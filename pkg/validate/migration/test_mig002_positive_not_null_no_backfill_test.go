//ff:func feature=validate type=test control=sequence topic=migration-safety
//ff:what MIG-002 positive — NOT NULL 승격 없이 backfill 도 없을 때 ERROR

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG002_Positive_NotNullNoBackfill(t *testing.T) {
	issues := []migration.SafetyIssue{
		{RuleID: "MIG-002", Level: migration.SafetyError, Message: "..."},
	}
	diags := Mig002NotNullWithoutBackfill(issues)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR, got %v", diags[0].Level)
	}
}
