//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-11 — belongs_to가 미정의 테이블을 참조하면 에러 진단 테스트
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

func TestFT11_BelongsToRef_NilTables(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: nil,
	}
	diags := ft11BelongsToRef(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags for nil tables, got %d", len(diags))
	}
}
