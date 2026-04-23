//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-02 positive — 모든 인자 키가 OpenAPI parameters 에 존재

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot02_AllKnown(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File: "page.tsx",
			Calls: []tsx.APICall{{
				OperationID: "getWorkflow", Line: 10,
				Args: []tsx.ArgBinding{{Key: "id"}, {Key: "version"}},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId":      {"getWorkflow": true},
		"OpenAPI.param.getWorkflow": {"id": true, "version": true},
	}})
	if diags := xot02ParameterMatch(fs); len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
