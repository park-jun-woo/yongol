//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what zerocov — xqs18CheckArg / xqs18CheckSeq 직접 호출 커버리지
package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs18CheckArg_EarlyReturns_ZeroCov(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{Type: "get", Model: "Item.FindByID", Line: 1}
	oapiParams := map[string]string{"item_id": "int64"}
	sqlcParams := map[string]bool{"ItemID": true}
	ddlColType := map[string]map[string]string{"items": {"item_id": "int32"}}

	cases := []struct {
		name string
		arg  ssacparser.Arg
	}{
		{"literal-source", ssacparser.Arg{Source: "", Field: "ItemID"}},
		{"var-source", ssacparser.Arg{Source: "tmpvar", Field: "ItemID"}},
		{"empty-field", ssacparser.Arg{Source: "request", Field: ""}},
		{"unknown-oapi", ssacparser.Arg{Source: "request", Field: "Unknown"}},
	}
	for _, c := range cases {
		if _, ok := xqs18CheckArg(fn, seq, c.arg, oapiParams, sqlcParams, true, ddlColType, "items"); ok {
			t.Errorf("%s: expected no diagnostic", c.name)
		}
	}

	// Missing sqlc param (hasSqlc true but field absent) → no diag.
	if _, ok := xqs18CheckArg(fn, seq, ssacparser.Arg{Source: "request", Field: "ItemID"},
		oapiParams, map[string]bool{}, true, ddlColType, "items"); ok {
		t.Error("expected no diagnostic when sqlc param absent")
	}

	// Compatible types (int64 OpenAPI vs int64 DDL) → no diag.
	okCol := map[string]map[string]string{"items": {"item_id": "int64"}}
	if _, ok := xqs18CheckArg(fn, seq, ssacparser.Arg{Source: "request", Field: "ItemID"},
		oapiParams, sqlcParams, true, okCol, "items"); ok {
		t.Error("expected no diagnostic for compatible int64 vs int64")
	}

	// Missing DDL column → no diag.
	if _, ok := xqs18CheckArg(fn, seq, ssacparser.Arg{Source: "request", Field: "ItemID"},
		oapiParams, sqlcParams, true, map[string]map[string]string{}, "items"); ok {
		t.Error("expected no diagnostic when DDL column missing")
	}
}
