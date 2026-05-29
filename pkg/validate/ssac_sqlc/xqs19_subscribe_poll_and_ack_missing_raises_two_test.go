//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs19_Subscribe_Poll_And_Ack_Missing_Raises_Two — @subscribe QueuePoll+QueueAck 부재 → 2건

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs19_Subscribe_Poll_And_Ack_Missing_Raises_Two(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "OnOrder", FileName: "svc.ssac",
			Subscribe: &ssacparser.SubscribeInfo{Topic: "order.created"},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"queue": queueInterface()},
	}
	diags := xqs19SsacBuiltinQueryRequired(fs)
	// Two ports (QueuePoll + QueueAck) both use Subscribe; both are missing.
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics (QueuePoll + QueueAck), got %d: %+v", len(diags), diags)
	}
}
