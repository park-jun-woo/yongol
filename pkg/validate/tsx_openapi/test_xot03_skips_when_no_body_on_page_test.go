//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-03 — 페이지 호출에 body 가 없으면 규칙을 발화하지 않음

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
