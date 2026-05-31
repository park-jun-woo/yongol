//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT13_FeatureTableRef_NoFire

package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
