//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-03 positive — useForm 필드가 request body schema 와 일치

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
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
