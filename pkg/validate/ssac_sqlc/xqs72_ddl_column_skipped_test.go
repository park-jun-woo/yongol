//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs72_DDLColumn_Skipped — Input key that maps to DDL column is skipped (XQS-18 scope)

package ssac_sqlc

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs72_DDLColumn_Skipped(t *testing.T) {
	fn := ssacparser.ServiceFunc{Name: "ListItems", FileName: "item.ssac"}
	seq := ssacparser.Sequence{
		Type:  "get",
		Model: "Item.ListByOrgPaged",
		Line:  15,
		Inputs: map[string]string{
			"OrgId": "request.org_id",
		},
	}
	oapiParams := map[string]string{"org_id": "int64"}
	sqlcParams := map[string]bool{"OrgId": true}
	ddlColType := map[string]map[string]string{
		"items": {"org_id": "int32"},
	}
	queryBody := "SELECT * FROM items WHERE org_id = sqlc.arg(org_id)"

	_, ok := xqs72CheckEntry(fn, seq, "OrgId", "request.org_id", oapiParams, sqlcParams, true, ddlColType, "items", queryBody)
	if ok {
		t.Fatal("expected no diagnostic when key maps to DDL column (XQS-18 scope)")
	}
}
