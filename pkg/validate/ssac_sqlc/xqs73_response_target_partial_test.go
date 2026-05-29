//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS73_ResponseTargetPartialSelect — @response target + 부분 SELECT 변수 에러 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXQS73_ResponseTargetPartialSelect verifies that @response target
// with a partial SELECT variable raises an error.
func TestXQS73_ResponseTargetPartialSelect(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "GetItem", FileName: "getitem.ssac",
			Sequences: []ssacparser.Sequence{
				{
					Type:   "get",
					Model:  "Item.FindSummary",
					Result: &ssacparser.Result{Type: "Item", Var: "item"},
					Line:   5,
				},
				{
					Type:   "response",
					Target: "item",
					Line:   10,
				},
			},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{{
			Name:        "ItemFindSummary",
			Model:       "Item",
			Method:      "FindSummary",
			Cardinality: "one",
			RowType:     "ItemFindSummaryRow",
			Body:        "SELECT id, title FROM items WHERE id = @id",
			SelectStar:  false,
			SelectCols:  []string{"id", "title"},
			File:        "items.sql",
			Line:        1,
		}},
	}
	diags := xqs73PartialSelectField(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "partial SELECT") {
		t.Fatalf("expected partial SELECT warning, got: %s", diags[0].Message)
	}
}
