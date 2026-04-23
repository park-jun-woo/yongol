//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-01 positive 테스트 — apiClient.<op>() 가 OpenAPI operationId 집합에 존재

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot01_Match(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File:  "page.tsx",
			Calls: []tsx.APICall{{OperationID: "listWorkflows", Line: 5}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId": {"listWorkflows": true},
	}})
	if diags := xot01OperationID(fs); len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
