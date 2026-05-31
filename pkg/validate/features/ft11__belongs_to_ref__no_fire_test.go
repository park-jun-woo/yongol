//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT11_BelongsToRef_NoFire

package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT11_BelongsToRef_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
			"tasks":    {BelongsTo: []string{"projects"}},
		},
	}
	diags := ft11BelongsToRef(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
