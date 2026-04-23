//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-01 negative 테스트 — operationId 가 없으면 ERROR

package tsx_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot01_MissingOperation(t *testing.T) {
	fs := &yongol.Fullstack{
		TSXPages: []tsx.PageSpec{{
			File:  "page.tsx",
			Calls: []tsx.APICall{{OperationID: "bogusOp", Line: 7}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId": {"listWorkflows": true},
	}})
	diags := xot01OperationID(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOT-1]") || !strings.Contains(diags[0].Message, "bogusOp") {
		t.Errorf("unexpected message: %q", diags[0].Message)
	}
}
