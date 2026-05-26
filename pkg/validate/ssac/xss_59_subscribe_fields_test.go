//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-59 fire/pass — @subscribe message field mismatch fires, all matched passes

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss59SubscribeFields(t *testing.T) {
	pub := parsessac.ServiceFunc{
		Name: "Pub", FileName: "s/pub.ssac",
		Sequences: []parsessac.Sequence{{Type: "publish", Topic: "order.completed", Inputs: map[string]string{"OrderID": "order.ID"}, Line: 5}},
	}
	sub := parsessac.ServiceFunc{
		Name: "Sub", FileName: "s/sub.ssac", Line: 3,
		Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
		Param:     &parsessac.ParamInfo{TypeName: "Msg", VarName: "message"},
		Structs:   []parsessac.StructInfo{{Name: "Msg", Fields: []parsessac.StructField{{Name: "OrderID", Type: "int64"}, {Name: "UserID", Type: "int64"}}}},
	}
	t.Run("fires_on_missing_field", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{pub, sub}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XSS-59]") || !strings.Contains(diags[0].Message, "UserID") {
			t.Errorf("unexpected message: %q", diags[0].Message)
		}
	})
	t.Run("passes_on_all_matched", func(t *testing.T) {
		pubAll := pub
		pubAll.Sequences = []parsessac.Sequence{{Type: "publish", Topic: "order.completed", Inputs: map[string]string{"OrderID": "order.ID", "UserID": "u.ID"}, Line: 5}}
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{pubAll, sub}}
		diags := xss59SubscribeFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
