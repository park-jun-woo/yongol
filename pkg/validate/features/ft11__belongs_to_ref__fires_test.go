//ff:func feature=validate type=test control=iteration dimension=1 topic=features-structural
//ff:what TestFT11_BelongsToRef_Fires

package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT11_BelongsToRef_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"tasks": {BelongsTo: []string{"projects", "ghost"}},
		},
	}
	diags := ft11BelongsToRef(fs)
	// "projects" and "ghost" are both undefined, so 2 diags.
	if len(diags) != 2 {
		t.Fatalf("want 2 diags, got %d", len(diags))
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, "[FT-11]") {
			t.Errorf("want [FT-11] prefix, got %s", d.Message)
		}
	}
}
