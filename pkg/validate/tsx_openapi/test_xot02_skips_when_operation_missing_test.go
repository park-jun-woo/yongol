//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-02 — operationId 자체가 없으면 XOT-1 이 커버하므로 skip

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot02_SkipsWhenOperationMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File:  "page.tsx",
			Calls: []tsx.APICall{{OperationID: "bogus", Args: []tsx.ArgBinding{{Key: "x"}}}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId": {"realOp": true},
	}})
	if diags := xot02ParameterMatch(fs); len(diags) != 0 {
		t.Errorf("XOT-1 covers unknown operation; XOT-2 should skip, got %+v", diags)
	}
}
