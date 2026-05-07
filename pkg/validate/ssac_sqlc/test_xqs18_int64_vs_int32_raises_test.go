//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs18_Int64VsInt32_Raises — OpenAPI int64 ↔ sqlc int32 → ERROR

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs18_Int64VsInt32_Raises(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{Type: "get", Model: "Item.FindByID", Line: 10}
	oapiParams := map[string]string{"item_id": "int64"}
	sqlcParams := map[string]bool{"ItemId": true}
	ddlColType := map[string]map[string]string{
		"items": {"item_id": "int32"},
	}

	d, ok := xqs18CheckInput(fn, seq, "ItemId", "request.item_id", oapiParams, sqlcParams, true, ddlColType, "items")
	if !ok {
		t.Fatal("expected diagnostic for int64 ↔ int32 mismatch, got none")
	}
	if !strings.Contains(d.Message, "XQS-18") {
		t.Errorf("diagnostic missing XQS-18: %s", d.Message)
	}
}
