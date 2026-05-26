//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-36 — stale response 통합 검증 (@get→@put→@response stale 검출, @get→@response 정상)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS36StaleResponse(t *testing.T) {
	t.Run("Fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Result: &ssac.Result{Type: "Order", Var: "order"}},
					{Type: "put", Line: 5, Model: "Order.Update"},
					{Type: "response", Line: 7, Fields: map[string]string{"data": "order.Name"}},
				}},
			},
		}
		diags := s36StaleResponse(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-36]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-36 diagnostic, got %d diags", len(diags))
		}
	})
	t.Run("Passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Result: &ssac.Result{Type: "Order", Var: "order"}},
					{Type: "response", Line: 7, Fields: map[string]string{"data": "order.Name"}},
				}},
			},
		}
		diags := s36StaleResponse(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-36]") {
				t.Errorf("unexpected S-36 diagnostic: %s", d.Message)
			}
		}
	})
	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s36StaleResponse(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
