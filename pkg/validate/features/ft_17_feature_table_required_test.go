//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-17 — feature table 필드 필수 검증 테스트
package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT17_FeatureTableRequired_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: []featparser.Feature{
			{Op: "Health", Path: "GET /health", Desc: "Health check", Table: "", Line: 3},
		},
	}
	diags := ft17FeatureTableRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-17]") {
		t.Errorf("want [FT-17] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "Health") {
		t.Errorf("want 'Health' in message, got %s", diags[0].Message)
	}
}

func TestFT17_FeatureTableRequired_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: []featparser.Feature{
			{Op: "ListProjects", Path: "GET /projects", Desc: "List", Table: "projects", Line: 5},
		},
	}
	diags := ft17FeatureTableRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}

func TestFT17_FeatureTableRequired_NilFeatures(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: nil,
	}
	diags := ft17FeatureTableRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
