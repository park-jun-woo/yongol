//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs72_Int64VsInt32_Raises — OpenAPI int64 ↔ sqlc int32 → ERROR

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs72_Int64VsInt32_Raises(t *testing.T) {
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
	queryBody := "SELECT * FROM items LIMIT sqlc.arg(per_page)"

	d, ok := xqs72CheckEntry(fn, seq, "PerPage", "request.per_page", oapiParams, sqlcParams, true, ddlColType, "items", queryBody)
	if !ok {
		t.Fatal("expected diagnostic for int64 ↔ int32 mismatch, got none")
	}
	if !strings.Contains(d.Message, "XQS-72") {
		t.Errorf("diagnostic missing XQS-72: %s", d.Message)
	}
	if !strings.Contains(d.Message, "int64") || !strings.Contains(d.Message, "int32") {
		t.Errorf("diagnostic should mention both widths: %s", d.Message)
	}
}
