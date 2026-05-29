//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs72_Int32Match_Passes — OpenAPI int32 ↔ sqlc int32 → no error

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs72_Int32Match_Passes(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "ListItems", FileName: "item.ssac"}
	seq := ssacparser.Sequence{
		Type:  "get",
		Model: "Item.ListByOrgPaged",
		Line:  15,
		Inputs: map[string]string{
			"PerPage": "request.per_page",
		},
	}
	oapiParams := map[string]string{"per_page": "int32"}
	sqlcParams := map[string]bool{"PerPage": true}
	ddlColType := map[string]map[string]string{
		"items": {"name": "string"},
	}
	queryBody := "SELECT * FROM items LIMIT sqlc.arg(per_page)"

	_, ok := xqs72CheckEntry(fn, seq, "PerPage", "request.per_page", oapiParams, sqlcParams, true, ddlColType, "items", queryBody)
	if ok {
		t.Fatal("expected no diagnostic for int32 ↔ int32 match")
	}
}
