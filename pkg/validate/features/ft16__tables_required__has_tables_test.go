//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT16_TablesRequired_HasTables

package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT16_TablesRequired_HasTables(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
	}
	diags := ft16TablesRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
