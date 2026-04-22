//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-01 테스트 — apiClient.<op>() 가 OpenAPI operationId 집합에 있는지

package tsx_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
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

func TestXot01_SkipWhenNoTSX(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId": {"any": true},
	}})
	if diags := xot01OperationID(fs); len(diags) != 0 {
		t.Errorf("expected no diagnostics when no TSX pages, got %+v", diags)
	}
}
