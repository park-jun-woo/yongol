//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-02 negative — 알 수 없는 인자 키는 ERROR

package tsx_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot02_UnknownKey(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File: "page.tsx",
			Calls: []tsx.APICall{{
				OperationID: "getWorkflow", Line: 10,
				Args: []tsx.ArgBinding{{Key: "id"}, {Key: "typo_name"}},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId":      {"getWorkflow": true},
		"OpenAPI.param.getWorkflow": {"id": true},
	}})
	diags := xot02ParameterMatch(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOT-2]") || !strings.Contains(diags[0].Message, "typo_name") {
		t.Errorf("unexpected message: %q", diags[0].Message)
	}
}
