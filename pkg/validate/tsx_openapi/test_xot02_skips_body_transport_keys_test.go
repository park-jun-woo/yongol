//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-02 — body/data 등 transport key 는 규칙에서 제외

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot02_SkipsBodyTransportKeys(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File: "page.tsx",
			Calls: []tsx.APICall{{
				OperationID: "createWorkflow",
				Args:        []tsx.ArgBinding{{Key: "body"}, {Key: "data"}},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId":         {"createWorkflow": true},
		"OpenAPI.param.createWorkflow": {},
	}})
	if diags := xot02ParameterMatch(fs); len(diags) != 0 {
		t.Errorf("body/data transport keys should be skipped, got %+v", diags)
	}
}
