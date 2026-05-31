//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT10_HasManyRef_NoFire

package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT10_HasManyRef_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {HasMany: []string{"tasks"}},
			"tasks":    {},
		},
	}
	diags := ft10HasManyRef(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
