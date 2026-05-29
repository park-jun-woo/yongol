//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xso16OpIDToFunc — nil ground/매칭/누락 SSaC func 검증

package openapi_ssac

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXso16OpIDToFunc_Unit(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xso16OpIDToFunc(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching func passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}},
		}
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"OpenAPI.operationId": {"getUser": true},
				"SSaC.funcName":       {"getUser": true},
			},
		}
		fs.SetGround(g)
		diags := xso16OpIDToFunc(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing func raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPILines: &oapiparser.LineIndex{Operations: map[string]int{}},
		}
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"OpenAPI.operationId": {"getUser": true},
				"SSaC.funcName":       {},
			},
		}
		fs.SetGround(g)
		diags := xso16OpIDToFunc(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XSO-16") {
			t.Errorf("Message missing XSO-16: %s", diags[0].Message)
		}
	})
}

func TestXso16OpIDToFunc(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
