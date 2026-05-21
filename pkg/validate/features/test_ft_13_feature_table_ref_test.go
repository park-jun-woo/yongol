//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-13 — feature의 table이 tables에 없으면 에러 진단 테스트
package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT13_FeatureTableRef_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: []featparser.Feature{
			{Op: "CreateTask", Path: "POST /tasks", Desc: "Create task", Table: "tasks", Line: 5},
		},
	}
	diags := ft13FeatureTableRef(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-13]") {
		t.Errorf("want [FT-13] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "tasks") {
		t.Errorf("want 'tasks' in message, got %s", diags[0].Message)
	}
}

func TestFT13_FeatureTableRef_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
			"tasks":    {},
		},
		Features: []featparser.Feature{
			{Op: "CreateTask", Path: "POST /tasks", Desc: "Create task", Table: "tasks", Line: 5},
			{Op: "ListProjects", Path: "GET /projects", Desc: "List projects", Table: "projects", Line: 8},
		},
	}
	diags := ft13FeatureTableRef(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}

func TestFT13_FeatureTableRef_NilFeatures(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]featparser.TableDef{
			"projects": {},
		},
		Features: nil,
	}
	diags := ft13FeatureTableRef(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags for nil features, got %d", len(diags))
	}
}
