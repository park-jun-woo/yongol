//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildResult — diagnostic → Result 변환 (ruleID 추출/level/locations/ruleIndex)
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestBuildResult_WithCatalog covers the full conversion: rule id extracted,
// level mapped, locations populated, and ruleIndex attached from the catalog.
func TestBuildResult_WithCatalog(t *testing.T) {
	cat := testCatalog() // S-1, S-2, X-3
	d := diagnostic.Diagnostic{
		File:    "specs/auth/login.ssac",
		Line:    15,
		Level:   diagnostic.LevelError,
		Message: "[S-2] foo not declared",
	}
	res, ruleID := buildResult(d, "specs", "/abs/specs", cat)

	if ruleID != "S-2" {
		t.Errorf("ruleID: got %q, want S-2", ruleID)
	}
	if res.RuleID != "S-2" {
		t.Errorf("res.RuleID: got %q, want S-2", res.RuleID)
	}
	if res.Level != "error" {
		t.Errorf("res.Level: got %q, want error", res.Level)
	}
	if res.Message.Text != "foo not declared" {
		t.Errorf("res.Message.Text: got %q", res.Message.Text)
	}
	if len(res.Locations) != 1 {
		t.Fatalf("locations: got %d, want 1", len(res.Locations))
	}
	if res.Locations[0].PhysicalLocation.ArtifactLocation.URI != "auth/login.ssac" {
		t.Errorf("uri: got %q", res.Locations[0].PhysicalLocation.ArtifactLocation.URI)
	}
	if res.RuleIndex == nil || *res.RuleIndex != 1 {
		t.Errorf("ruleIndex: want 1, got %v", res.RuleIndex)
	}
}

// TestBuildResult_NoRuleNoCatalog covers a prefix-less message with no catalog:
// empty ruleID, no ruleIndex, and (no file) no locations.
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
