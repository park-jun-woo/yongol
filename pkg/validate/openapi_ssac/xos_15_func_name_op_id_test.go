//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos15FuncNameOpID — nil ground/매칭/누락 operationId 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos15FuncNameOpID_Unit(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos15FuncNameOpID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching operationId passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", FileName: "user.ssac"},
			},
		}
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"OpenAPI.operationId": {"getUser": true},
			},
		}
		fs.SetGround(g)
		diags := xos15FuncNameOpID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing operationId raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", FileName: "user.ssac"},
			},
		}
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"OpenAPI.operationId": {},
			},
		}
		fs.SetGround(g)
		diags := xos15FuncNameOpID(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-15") {
			t.Errorf("Message missing XOS-15: %s", diags[0].Message)
		}
	})

	t.Run("subscribe func skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "onOrderCreated", FileName: "order.ssac", Subscribe: &ssac.SubscribeInfo{}},
			},
		}
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"OpenAPI.operationId": {},
			},
		}
		fs.SetGround(g)
		diags := xos15FuncNameOpID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
