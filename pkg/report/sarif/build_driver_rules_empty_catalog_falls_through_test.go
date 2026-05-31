//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildDriverRules — catalog 전체 / fired-only fallback / 빈 fired nil 분기 검증
package sarif

import (
	"testing"

	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

func TestBuildDriverRules_EmptyCatalogFallsThrough(t *testing.T) {
	emptyCat := rulecatalog.NewCatalog(nil)
	got := buildDriverRules(emptyCat, map[string]struct{}{"S-9": {}})
	if len(got) != 1 || got[0].ID != "S-9" {
		t.Errorf("empty catalog should fall back to fired: got %+v", got)
	}
}
