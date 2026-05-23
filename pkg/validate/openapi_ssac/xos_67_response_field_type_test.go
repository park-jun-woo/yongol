//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos67ResponseFieldType — nil ground/subscribe skip/non-response skip 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos67ResponseFieldType_Unit(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("subscribe func skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "onOrder", Subscribe: &ssac.SubscribeInfo{}},
			},
		}
		g := &rule.Ground{Types: map[string]string{}}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("non-response sequence skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{{Type: "get"}}},
			},
		}
		g := &rule.Ground{Types: map[string]string{}}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
