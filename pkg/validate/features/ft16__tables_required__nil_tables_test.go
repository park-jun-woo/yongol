//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT16_TablesRequired_NilTables

package features

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT16_TablesRequired_NilTables(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: nil,
	}
	diags := ft16TablesRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-16]") {
		t.Errorf("want [FT-16] prefix, got %s", diags[0].Message)
	}
}
