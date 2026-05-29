//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT17_FeatureTableRequired_NilFeatures

package features

import (
	"testing"
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT17_FeatureTableRequired_NilFeatures(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: nil,
	}
	diags := ft17FeatureTableRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
