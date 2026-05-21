//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-16 — tables 섹션 필수 검증 테스트
package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
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

func TestFT16_TablesRequired_HasTables(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
	}
	diags := ft16TablesRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
