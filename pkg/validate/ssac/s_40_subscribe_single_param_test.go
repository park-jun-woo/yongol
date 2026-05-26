//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-40 — @subscribe 파라미터는 단일 'message' 변수여야 한다

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS40SubscribeSingleParam(t *testing.T) {
	t.Run("Fires_nil_param", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Line: 1, Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Param: nil},
			},
		}
		diags := s40SubscribeSingleParam(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-40]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-40 diagnostic for nil param")
		}
	})

	t.Run("Fires_wrong_name", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Line: 1, Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Param: &ssac.ParamInfo{TypeName: "Msg", VarName: "msg"}},
			},
		}
		diags := s40SubscribeSingleParam(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-40]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-40 diagnostic for wrong param name 'msg'")
		}
	})

	t.Run("Passes_message", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "worker.ssac", Line: 1, Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Param: &ssac.ParamInfo{TypeName: "Msg", VarName: "message"}},
			},
		}
		diags := s40SubscribeSingleParam(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-40]") {
				t.Errorf("unexpected S-40 diagnostic: %s", d.Message)
			}
		}
	})

	t.Run("Skips_http_func", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Line: 1, Subscribe: nil},
			},
		}
		diags := s40SubscribeSingleParam(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for HTTP func", len(diags))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s40SubscribeSingleParam(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
