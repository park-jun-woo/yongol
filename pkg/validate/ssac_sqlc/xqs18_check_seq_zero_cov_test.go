//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what zerocov — xqs18CheckArg / xqs18CheckSeq 직접 호출 커버리지
package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs18CheckSeq_ZeroCov(t *testing.T) {
	oapiParams := map[string]string{"item_id": "int64"}
	paramMap := map[string]map[string]bool{}
	ddlColType := map[string]map[string]string{"items": {"item_id": "int32"}}

	// type == "call" → skipped (nil).
	if diags := xqs18CheckSeq(ssacparser.ServiceFunc{}, ssacparser.Sequence{Type: "call"},
		oapiParams, paramMap, ddlColType); diags != nil {
		t.Errorf("call seq should be skipped, got %v", diags)
	}

	// empty Model → skipped (nil).
	if diags := xqs18CheckSeq(ssacparser.ServiceFunc{}, ssacparser.Sequence{Type: "get", Model: ""},
		oapiParams, paramMap, ddlColType); diags != nil {
		t.Errorf("empty-model seq should be skipped, got %v", diags)
	}

	// Full sequence with a mismatching Arg + mismatching Input → 2 diagnostics.
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{
		Type:  "get",
		Model: "Item.FindByID",
		Line:  20,
		Args:  []ssacparser.Arg{{Source: "request", Field: "ItemID"}},
		Inputs: map[string]string{
			"ItemID": "request.item_id",
		},
	}
	// queryName resolves to a sqlc query that lists ItemID as a param.
	qn := resolveQueryName(seq)
	paramMap[qn] = map[string]bool{"ItemID": true}

	diags := xqs18CheckSeq(fn, seq, oapiParams, paramMap, ddlColType)
	if len(diags) == 0 {
		t.Fatal("expected at least one diagnostic from Args/Inputs iteration")
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, "XQS-18") {
			t.Errorf("diagnostic missing XQS-18: %s", d.Message)
		}
	}
}
