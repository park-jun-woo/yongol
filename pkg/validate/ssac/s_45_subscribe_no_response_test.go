//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-45 — @subscribe 에서 @response 사용 금지 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS45SubscribeNoResponse(t *testing.T) {
	t.Run("Fires_response_in_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID"},
					{Type: "response", Line: 5, Fields: map[string]string{"data": "order"}},
				}},
			},
		}
		diags := s45SubscribeNoResponse(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-45]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-45 diagnostic for @response in subscribe")
		}
	})

	t.Run("Passes_no_response", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID"},
					{Type: "put", Line: 5, Model: "Order.Update"},
				}},
			},
		}
		diags := s45SubscribeNoResponse(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-45]") {
				t.Errorf("unexpected S-45 diagnostic: %s", d.Message)
			}
		}
	})

	t.Run("Skips_http_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Subscribe: nil, Sequences: []ssac.Sequence{
					{Type: "response", Line: 5, Fields: map[string]string{"data": "order"}},
				}},
			},
		}
		diags := s45SubscribeNoResponse(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for HTTP func", len(diags))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s45SubscribeNoResponse(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
