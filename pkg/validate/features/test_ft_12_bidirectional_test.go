//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-12 — has_many에 대응하는 belongs_to 없으면 WARN 테스트
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
