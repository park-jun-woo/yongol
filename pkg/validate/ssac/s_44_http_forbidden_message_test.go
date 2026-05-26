//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-44 — HTTP 핸들러에서 message 사용 금지 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS44HTTPForbiddenMessage(t *testing.T) {
	t.Run("Fires_message_in_http", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Subscribe: nil, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "message", Field: "OrderID"}}},
				}},
			},
		}
		diags := s44HTTPForbiddenMessage(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-44]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-44 diagnostic for message in HTTP func")
		}
	})

	t.Run("Passes_request_in_http", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Subscribe: nil, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "request", Field: "ID"}}},
				}},
			},
		}
		diags := s44HTTPForbiddenMessage(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-44]") {
				t.Errorf("unexpected S-44 diagnostic for request source: %s", d.Message)
			}
		}
	})

	t.Run("Skips_subscribe_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Args: []ssac.Arg{{Source: "message", Field: "OrderID"}}},
				}},
			},
		}
		diags := s44HTTPForbiddenMessage(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for subscribe func", len(diags))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s44HTTPForbiddenMessage(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
