//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestXqs18_NoFormatLoose_Passes — format 미명시 ↔ int32/int64 → PASS (하위 호환)

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs18_NoFormatLoose_Passes(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{Type: "get", Model: "Item.FindByID", Line: 10}
	sqlcParams := map[string]bool{"ItemId": true}

	for _, goType := range []string{"int32", "int64"} {
		oapiParams := map[string]string{"item_id": "integer"}
		ddlColType := map[string]map[string]string{
			"items": {"item_id": goType},
		}

		_, ok := xqs18CheckInput(fn, seq, "ItemId", "request.item_id", oapiParams, sqlcParams, true, ddlColType, "items")
		if ok {
			t.Errorf("expected no diagnostic for 'integer' (no format) ↔ %q, got one", goType)
		}
	}
}
