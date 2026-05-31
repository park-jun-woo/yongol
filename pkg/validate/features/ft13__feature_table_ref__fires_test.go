//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT13_FeatureTableRef_Fires

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
