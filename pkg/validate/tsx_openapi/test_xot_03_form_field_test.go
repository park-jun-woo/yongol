//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-03 테스트 — useForm register('x') ↔ request body schema 일치 (WARNING)

package tsx_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestXot03_AllFieldsKnown(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File:       "page.tsx",
			Calls:      []tsx.APICall{{OperationID: "createWorkflow"}},
			FormFields: []tsx.FormField{{Name: "title"}, {Name: "trigger_event"}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.request.createWorkflow": {"title": true, "trigger_event": true, "description": true},
	}})
	if diags := xot03FormField(fs); len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}

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

func TestXot03_SkipsWhenNoBodyOnPage(t *testing.T) {
	// No apiClient calls have a request body — rule must not fire to avoid
	// false positives (helper-based forms).
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File:       "page.tsx",
			Calls:      []tsx.APICall{{OperationID: "listWorkflows"}}, // GET
			FormFields: []tsx.FormField{{Name: "search"}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
	if diags := xot03FormField(fs); len(diags) != 0 {
		t.Errorf("want 0 diagnostics when no body on page, got %+v", diags)
	}
}
