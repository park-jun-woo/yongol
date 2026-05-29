//ff:func feature=validate type=test control=sequence topic=migration-safety
//ff:what MIG-005 positive — 타입 변환에 USING CAST 가 없으면 WARNING

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG005_Positive_CastMissing(t *testing.T) {
	issues := []migration.SafetyIssue{
		{RuleID: "MIG-005", Level: migration.SafetyWarning, Message: "..."},
	}
	diags := Mig005CastMissing(issues)
	if len(diags) != 1 || diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("expected 1 WARN, got %+v", diags)
	}
}
