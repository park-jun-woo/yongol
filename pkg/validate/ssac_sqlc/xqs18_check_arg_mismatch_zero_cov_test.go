//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zerocov — xqs18CheckArg / xqs18CheckSeq 직접 호출 커버리지
package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs18CheckArg_Mismatch_ZeroCov(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{Type: "get", Model: "Item.FindByID", Line: 12}
	arg := ssacparser.Arg{Source: "request", Field: "ItemID"}
	oapiParams := map[string]string{"item_id": "int64"}
	sqlcParams := map[string]bool{"ItemID": true}
	ddlColType := map[string]map[string]string{
		"items": {"item_id": "int32"},
	}

	d, ok := xqs18CheckArg(fn, seq, arg, oapiParams, sqlcParams, true, ddlColType, "items")
	if !ok {
		t.Fatal("expected diagnostic for int64 vs int32 mismatch")
	}
	if !strings.Contains(d.Message, "XQS-18") {
		t.Errorf("diagnostic missing XQS-18: %s", d.Message)
	}
}
