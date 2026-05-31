//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestRuleFromMeta — RuleMeta → SARIF Rule 변환 (full / ERROR-min / WARNING / unknown-level)
package sarif

import (
	"testing"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

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
