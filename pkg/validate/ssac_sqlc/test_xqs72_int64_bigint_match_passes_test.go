//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs72_Int64BigintMatch_Passes — OpenAPI int64 ↔ sqlc ::bigint → no error

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs72_Int64BigintMatch_Passes(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "ListItems", FileName: "item.ssac"}
	seq := ssacparser.Sequence{
		Type:  "get",
		Model: "Item.ListByOrgPaged",
		Line:  15,
		Inputs: map[string]string{
			"PerPage": "request.per_page",
		},
	}
	oapiParams := map[string]string{"per_page": "int64"}
	sqlcParams := map[string]bool{"PerPage": true}
	ddlColType := map[string]map[string]string{
		"items": {"name": "string"},
	}
	queryBody := "SELECT * FROM items LIMIT sqlc.arg(per_page)::bigint"

	_, ok := xqs72CheckEntry(fn, seq, "PerPage", "request.per_page", oapiParams, sqlcParams, true, ddlColType, "items", queryBody)
	if ok {
		t.Fatal("expected no diagnostic for int64 ↔ int64 (::bigint) match")
	}
}
