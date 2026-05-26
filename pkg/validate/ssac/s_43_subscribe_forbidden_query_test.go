//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-43 — @subscribe 에서 query 사용 금지 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS43SubscribeForbiddenQuery(t *testing.T) {
	t.Run("Fires_query_in_args", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.List", Args: []ssac.Arg{{Source: "query", Field: "Page"}}},
				}},
			},
		}
		diags := s43SubscribeForbiddenQuery(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-43]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-43 diagnostic for query in subscribe args")
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
		diags := s43SubscribeForbiddenQuery(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-43]") {
				t.Errorf("unexpected S-43 diagnostic for message source: %s", d.Message)
			}
		}
	})

	t.Run("Skips_http_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Subscribe: nil, Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.List", Args: []ssac.Arg{{Source: "query", Field: "Page"}}},
				}},
			},
		}
		diags := s43SubscribeForbiddenQuery(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for HTTP func", len(diags))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s43SubscribeForbiddenQuery(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
