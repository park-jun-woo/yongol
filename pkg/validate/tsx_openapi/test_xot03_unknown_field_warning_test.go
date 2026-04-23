//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-03 — body schema 에 없는 필드는 WARNING

package tsx_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot03_UnknownFieldWarning(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File:       "page.tsx",
			Calls:      []tsx.APICall{{OperationID: "createWorkflow"}},
			FormFields: []tsx.FormField{{Name: "title"}, {Name: "wrong_name", Line: 20}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.request.createWorkflow": {"title": true},
	}})
	diags := xot03FormField(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("expected WARNING, got %s", diags[0].Level)
	}
	if !strings.Contains(diags[0].Message, "[XOT-3]") || !strings.Contains(diags[0].Message, "wrong_name") {
		t.Errorf("unexpected message: %q", diags[0].Message)
	}
}
