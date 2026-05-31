//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT10_HasManyRef_Fires

package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT10_HasManyRef_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {HasMany: []string{"tasks", "unknown_table"}},
			"tasks":    {},
		},
	}
	diags := ft10HasManyRef(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-10]") {
		t.Errorf("want [FT-10] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "unknown_table") {
		t.Errorf("want unknown_table in message, got %s", diags[0].Message)
	}
}
