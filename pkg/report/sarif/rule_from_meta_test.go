//ff:func feature=report type=test control=selection topic=sarif
//ff:what TestRuleFromMeta — RuleMeta → SARIF Rule 변환 (full / ERROR-min / WARNING / unknown-level)
package sarif

import (
	"testing"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

// TestRuleFromMeta_Full covers every populated field plus the ERROR config
// branch and the source→properties mapping.
func TestRuleFromMeta_Full(t *testing.T) {
	m := rulecatalog.RuleMeta{
		ID:            "S-27",
		Level:         "ERROR",
		Description:   "variable not declared",
		Source:        "pkg/validate/ssac/s_27.go",
		SectionAnchor: "a-ssac-internal",
	}
	r := ruleFromMeta(m)
	if r.ID != "S-27" || r.Name != "S-27" {
		t.Errorf("id/name: %+v", r)
	}
	if r.ShortDescription == nil || r.ShortDescription.Text != "variable not declared" {
		t.Errorf("shortDescription: %+v", r.ShortDescription)
	}
	if r.HelpURI != helpURIBase+"#a-ssac-internal" {
		t.Errorf("helpUri: got %q", r.HelpURI)
	}
	if r.DefaultConfiguration == nil || r.DefaultConfiguration.Level != "error" {
		t.Errorf("defaultConfiguration: %+v", r.DefaultConfiguration)
	}
	if r.Properties == nil || r.Properties["source"] != "pkg/validate/ssac/s_27.go" {
		t.Errorf("properties: %+v", r.Properties)
	}
}

// TestRuleFromMeta_WarningMinimal covers the WARNING config branch and the
// empty-description / empty-anchor / empty-source branches (all omitted).
func TestRuleFromMeta_WarningMinimal(t *testing.T) {
	m := rulecatalog.RuleMeta{ID: "S-36", Level: "WARNING"}
	r := ruleFromMeta(m)
	if r.DefaultConfiguration == nil || r.DefaultConfiguration.Level != "warning" {
		t.Errorf("defaultConfiguration: %+v", r.DefaultConfiguration)
	}
	if r.ShortDescription != nil {
		t.Errorf("shortDescription should be nil, got %+v", r.ShortDescription)
	}
	if r.HelpURI != "" {
		t.Errorf("helpUri should be empty, got %q", r.HelpURI)
	}
	if r.Properties != nil {
		t.Errorf("properties should be nil, got %+v", r.Properties)
	}
}

// TestRuleFromMeta_UnknownLevel covers the default (no config) branch when the
// level is neither ERROR nor WARNING.
func TestRuleFromMeta_UnknownLevel(t *testing.T) {
	m := rulecatalog.RuleMeta{ID: "Q-1", Level: "INFO"}
	r := ruleFromMeta(m)
	if r.DefaultConfiguration != nil {
		t.Errorf("defaultConfiguration should be nil for unknown level, got %+v", r.DefaultConfiguration)
	}
}
