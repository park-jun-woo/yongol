//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zerocov — xqs18CheckArg / xqs18CheckSeq 직접 호출 커버리지

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestXqs18CheckArg_Mismatch_ZeroCov drives xqs18CheckArg through its mismatch
// path: a request-sourced Arg whose OpenAPI type (int64) is incompatible with
// the DDL column Go type (int32) must yield a diagnostic.
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

// TestXqs18CheckArg_EarlyReturns_ZeroCov drives the non-mismatch / early-return
// branches: non-request source, empty field, unknown OpenAPI param, missing
// sqlc param, missing DDL column, and a compatible type.
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

// TestXqs18CheckSeq_ZeroCov drives xqs18CheckSeq across its skip branches plus
// the full Args + Inputs iteration that aggregates xqs18CheckArg / CheckInput.
func TestXqs18CheckSeq_ZeroCov(t *testing.T) {
	oapiParams := map[string]string{"item_id": "int64"}
	paramMap := map[string]map[string]bool{}
	ddlColType := map[string]map[string]string{"items": {"item_id": "int32"}}

	// type == "call" → skipped (nil).
	if diags := xqs18CheckSeq(ssacparser.ServiceFunc{}, ssacparser.Sequence{Type: "call"},
		oapiParams, paramMap, ddlColType); diags != nil {
		t.Errorf("call seq should be skipped, got %v", diags)
	}

	// empty Model → skipped (nil).
	if diags := xqs18CheckSeq(ssacparser.ServiceFunc{}, ssacparser.Sequence{Type: "get", Model: ""},
		oapiParams, paramMap, ddlColType); diags != nil {
		t.Errorf("empty-model seq should be skipped, got %v", diags)
	}

	// Full sequence with a mismatching Arg + mismatching Input → 2 diagnostics.
	fn := ssacparser.ServiceFunc{Name: "GetItem", FileName: "item.ssac"}
	seq := ssacparser.Sequence{
		Type:  "get",
		Model: "Item.FindByID",
		Line:  20,
		Args:  []ssacparser.Arg{{Source: "request", Field: "ItemID"}},
		Inputs: map[string]string{
			"ItemID": "request.item_id",
		},
	}
	// queryName resolves to a sqlc query that lists ItemID as a param.
	qn := resolveQueryName(seq)
	paramMap[qn] = map[string]bool{"ItemID": true}

	diags := xqs18CheckSeq(fn, seq, oapiParams, paramMap, ddlColType)
	if len(diags) == 0 {
		t.Fatal("expected at least one diagnostic from Args/Inputs iteration")
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, "XQS-18") {
			t.Errorf("diagnostic missing XQS-18: %s", d.Message)
		}
	}
}
