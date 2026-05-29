//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT12_Bidirectional_ChildNotDefined

package features

import (
	"testing"
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT12_Bidirectional_ChildNotDefined(t *testing.T) {
	// When child table is not defined, FT-10 catches it. FT-12 should skip.
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {HasMany: []string{"ghost"}},
		},
	}
	diags := ft12Bidirectional(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (FT-10 catches undefined child), got %d", len(diags))
	}
}
