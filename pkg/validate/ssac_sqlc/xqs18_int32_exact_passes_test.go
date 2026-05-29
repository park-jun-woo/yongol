//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs18_Int32Exact_Passes — OpenAPI int32 ↔ sqlc int32 → PASS

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs18_Int32Exact_Passes(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{Type: "get", Model: "Item.FindByID", Line: 10}
	oapiParams := map[string]string{"item_id": "int32"}
	sqlcParams := map[string]bool{"ItemId": true}
	ddlColType := map[string]map[string]string{
		"items": {"item_id": "int32"},
	}

	_, ok := xqs18CheckInput(fn, seq, "ItemId", "request.item_id", oapiParams, sqlcParams, true, ddlColType, "items")
	if ok {
		t.Fatal("expected no diagnostic for int32 ↔ int32 exact match")
	}
}
