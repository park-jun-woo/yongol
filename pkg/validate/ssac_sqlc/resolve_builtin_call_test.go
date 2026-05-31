//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestResolveBuiltinCall(t *testing.T) {
	p, m := resolveBuiltinCall(ssacparser.Sequence{Type: "call", Package: "session", Model: "Get"}, false)
	if p != "session" || m != "Get" {
		t.Errorf("call = (%q,%q), want (session,Get)", p, m)
	}
	p, m = resolveBuiltinCall(ssacparser.Sequence{Type: "call", Package: "", Model: "Get"}, false)
	if p != "" || m != "" {
		t.Errorf("call missing pkg = (%q,%q), want empties", p, m)
	}
	p, m = resolveBuiltinCall(ssacparser.Sequence{Type: "publish"}, false)
	if p != "queue" || m != "Publish" {
		t.Errorf("publish = (%q,%q), want (queue,Publish)", p, m)
	}
	p, m = resolveBuiltinCall(ssacparser.Sequence{Type: "get"}, false)
	if p != "" || m != "" {
		t.Errorf("get = (%q,%q), want empties", p, m)
	}
}
