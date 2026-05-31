//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestRuleFromMeta — RuleMeta → SARIF Rule 변환 (full / ERROR-min / WARNING / unknown-level)
package sarif

import (
	"testing"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

func TestRuleFromMeta_UnknownLevel(t *testing.T) {
	m := rulecatalog.RuleMeta{ID: "Q-1", Level: "INFO"}
	r := ruleFromMeta(m)
	if r.DefaultConfiguration != nil {
		t.Errorf("defaultConfiguration should be nil for unknown level, got %+v", r.DefaultConfiguration)
	}
}
