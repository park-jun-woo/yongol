//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertResponse_ZeroCov(t *testing.T) {
	// fields branch (sorted)
	op := convertResponse(ssac.Sequence{Fields: map[string]string{"b": "y", "a": "x"}})
	if op.Kind != OpResponse || op.Response == nil {
		t.Fatalf("expected OpResponse, got %+v", op)
	}
	if len(op.Response.Fields) != 2 || op.Response.Fields[0].Name != "a" {
		t.Errorf("fields not sorted: %+v", op.Response.Fields)
	}
	// target branch
	op = convertResponse(ssac.Sequence{Target: "course"})
	if op.Response.SingleVar != "course" {
		t.Errorf("SingleVar = %q", op.Response.SingleVar)
	}
	// empty branch
	op = convertResponse(ssac.Sequence{})
	if op.Response.SingleVar != "" || len(op.Response.Fields) != 0 {
		t.Errorf("expected empty response, got %+v", op.Response)
	}
}
