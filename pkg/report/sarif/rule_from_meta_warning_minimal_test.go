//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestRuleFromMeta — RuleMeta → SARIF Rule 변환 (full / ERROR-min / WARNING / unknown-level)
package sarif

import (
	"testing"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

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
