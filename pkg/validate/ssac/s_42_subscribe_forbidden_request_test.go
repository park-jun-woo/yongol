//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-42 — @subscribe 에서 request 사용 금지 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS42SubscribeForbiddenRequest(t *testing.T) {
	t.Run("Fires_request_in_args", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "request", Field: "OrderID"}}},
				}},
			},
		}
		diags := s42SubscribeForbiddenRequest(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-42]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-42 diagnostic for request in subscribe args")
		}
	})

	t.Run("Passes_message_source", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "message", Field: "OrderID"}}},
				}},
			},
		}
		diags := s42SubscribeForbiddenRequest(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-42]") {
				t.Errorf("unexpected S-42 diagnostic for message source: %s", d.Message)
			}
		}
	})

	t.Run("Skips_http_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Subscribe: nil, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "request", Field: "ID"}}},
				}},
			},
		}
		diags := s42SubscribeForbiddenRequest(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for HTTP func", len(diags))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s42SubscribeForbiddenRequest(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
