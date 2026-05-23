//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT12_Bidirectional_NoFire

package features

import (
	"testing"
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT12_Bidirectional_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {HasMany: []string{"tasks"}},
			"tasks":    {BelongsTo: []string{"projects"}},
		},
	}
	diags := ft12Bidirectional(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
