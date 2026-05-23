//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT12_Bidirectional_Fires

package features

import (
	"strings"
	"testing"
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT12_Bidirectional_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {HasMany: []string{"tasks"}},
			"tasks":    {}, // no belongs_to
		},
	}
	diags := ft12Bidirectional(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-12]") {
		t.Errorf("want [FT-12] prefix, got %s", diags[0].Message)
	}
	if diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("want WARNING, got %s", diags[0].Level)
	}
}
