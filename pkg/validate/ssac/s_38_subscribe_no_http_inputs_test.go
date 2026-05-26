//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-38 — @subscribe 에서 HTTP-only 소스(request/query/currentUser) 사용 금지 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS38SubscribeNoHTTPInputs(t *testing.T) {
	sub := &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}
	t.Run("Fires_request", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "w.ssac", Subscribe: sub, Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "O.F", Inputs: map[string]string{"id": "request.OrderID"}},
		}}}}
		assertDiag(t, s38SubscribeNoHTTPInputs(fs), "[S-38]")
	})
	t.Run("Fires_query", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "w.ssac", Subscribe: sub, Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "O.L", Inputs: map[string]string{"p": "query.Page"}},
		}}}}
		assertDiag(t, s38SubscribeNoHTTPInputs(fs), "[S-38]")
	})
	t.Run("Fires_currentUser", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "w.ssac", Subscribe: sub, Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "U.F", Inputs: map[string]string{"id": "currentUser.ID"}},
		}}}}
		assertDiag(t, s38SubscribeNoHTTPInputs(fs), "[S-38]")
	})
	t.Run("Passes_message", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "w.ssac", Subscribe: sub, Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "O.F", Inputs: map[string]string{"id": "message.OrderID"}},
		}}}}
		assertNoDiag(t, s38SubscribeNoHTTPInputs(fs), "[S-38]")
	})
	t.Run("Passes_http", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Subscribe: nil, Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "O.F", Inputs: map[string]string{"id": "request.ID"}},
		}}}}
		assertNoDiag(t, s38SubscribeNoHTTPInputs(fs), "[S-38]")
	})
	t.Run("Empty", func(t *testing.T) {
		assertNoDiag(t, s38SubscribeNoHTTPInputs(&yongol.Fullstack{}), "[S-38]")
	})
}
