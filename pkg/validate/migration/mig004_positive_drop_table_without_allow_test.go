//ff:func feature=validate type=test control=sequence topic=migration-safety
//ff:what MIG-004 positive — DROP TABLE 에 @allow_destructive 없으면 WARNING

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG004_Positive_DropTableWithoutAllow(t *testing.T) {
	issues := []migration.SafetyIssue{
		{RuleID: "MIG-004", Level: migration.SafetyWarning, Message: "..."},
	}
	diags := Mig004DestructiveWithoutAllow(issues)
	if len(diags) != 1 || diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("expected 1 WARN, got %+v", diags)
	}
}
