//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT16_TablesRequired_EmptyMap

package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT16_TablesRequired_EmptyMap(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{},
	}
	diags := ft16TablesRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-16]") {
		t.Errorf("want [FT-16] prefix, got %s", diags[0].Message)
	}
}
