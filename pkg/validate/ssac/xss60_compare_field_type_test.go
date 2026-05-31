//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestSSaCXss60Helpers — unit tests for the pure ssac validate helper functions
package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60CompareFieldType(t *testing.T) {
	fn := parsessac.ServiceFunc{
		FileName:  "order.ssac",
		Line:      10,
		Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
	}
	// Type mismatch → diagnostic returned.
	pub := map[string]string{"OrderID": "int64"}
	diag, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "OrderID", Type: "string"}, pub)
	if !ok {
		t.Fatal("expected mismatch diagnostic")
	}
	if diag.Line != 10 || diag.File != "order.ssac" {
		t.Errorf("diag location = %s:%d", diag.File, diag.Line)
	}

	// Matching type → no diagnostic.
	if _, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "OrderID", Type: "int64"}, pub); ok {
		t.Error("matching type should not produce a diagnostic")
	}
	// Field not in publish map → no diagnostic.
	if _, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "Other", Type: "string"}, pub); ok {
		t.Error("absent field should not produce a diagnostic")
	}
	// Empty publish type → no diagnostic.
	if _, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "Empty", Type: "string"}, map[string]string{"Empty": ""}); ok {
		t.Error("empty publish type should not produce a diagnostic")
	}
}
