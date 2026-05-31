//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT13_FeatureTableRef_NilFeatures

package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT13_FeatureTableRef_NilFeatures(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: nil,
	}
	diags := ft13FeatureTableRef(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags for nil features, got %d", len(diags))
	}
}
