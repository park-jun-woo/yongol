//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildResult — diagnostic → Result 변환 (ruleID 추출/level/locations/ruleIndex)
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestBuildResult_NoRuleNoCatalog(t *testing.T) {
	d := diagnostic.Diagnostic{
		Level:   diagnostic.LevelWarning,
		Message: "plain warning",
	}
	res, ruleID := buildResult(d, "", "", nil)
	if ruleID != "" {
		t.Errorf("ruleID: got %q, want empty", ruleID)
	}
	if res.RuleID != "" {
		t.Errorf("res.RuleID: got %q, want empty", res.RuleID)
	}
	if res.Level != "warning" {
		t.Errorf("res.Level: got %q, want warning", res.Level)
	}
	if res.Locations != nil {
		t.Errorf("locations should be nil (no file), got %+v", res.Locations)
	}
	if res.RuleIndex != nil {
		t.Errorf("ruleIndex should be nil, got %v", *res.RuleIndex)
	}
}
