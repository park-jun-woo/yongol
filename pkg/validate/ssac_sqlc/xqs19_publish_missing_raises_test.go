//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs19_Publish_Missing_Raises — @publish QueuePublish 부재 → [XQS-19]

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs19_Publish_Missing_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Enqueue", FileName: "svc.ssac",
			Sequences: []ssacparser.Sequence{{Type: "publish", Topic: "order.created"}},
		}},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{"queue": queueInterface()},
	}
	diags := xqs19SsacBuiltinQueryRequired(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "QueuePublish") {
		t.Errorf("expected single XQS-19 naming QueuePublish, got: %+v", diags)
	}
}
