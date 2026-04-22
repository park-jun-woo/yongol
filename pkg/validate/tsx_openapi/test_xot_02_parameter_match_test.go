//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-02 테스트 — apiClient 호출 인자 키 ↔ OpenAPI parameters 일치 검증

package tsx_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
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
		"OpenAPI.operationId":     {"getWorkflow": true},
		"OpenAPI.param.getWorkflow": {"id": true, "version": true},
	}})
	if diags := xot02ParameterMatch(fs); len(diags) != 0 {
		t.Fatalf("want 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}

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
		"OpenAPI.operationId":       {"getWorkflow": true},
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
		"OpenAPI.operationId":          {"createWorkflow": true},
		"OpenAPI.param.createWorkflow": {},
	}})
	if diags := xot02ParameterMatch(fs); len(diags) != 0 {
		t.Errorf("body/data transport keys should be skipped, got %+v", diags)
	}
}

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
