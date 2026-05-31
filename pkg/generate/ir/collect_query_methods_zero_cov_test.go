//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgHelpersZeroCov — classifyFieldArgLocation / collectFieldArgSlices / opModelName / collectQueryMethods 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectQueryMethods_ZeroCov(t *testing.T) {
	seqs := []ssac.Sequence{
		{Type: ssac.SeqGet, Model: "Course.FindByID", Package: "queries"},
		{Type: ssac.SeqGet, Model: "Course.FindByID", Package: "queries"}, // dup → skipped
		{Type: ssac.SeqPost, Model: "Order.Insert", Package: "session"},
		{Type: ssac.SeqResponse, Model: "ignored"}, // non-CRUD → skipped
	}
	methods := collectQueryMethods(seqs)
	if len(methods) != 2 {
		t.Fatalf("expected 2 query methods, got %d: %+v", len(methods), methods)
	}
	if methods[0].Name != "CourseFindByID" {
		t.Errorf("method[0].Name = %q, want CourseFindByID", methods[0].Name)
	}
	if methods[1].Package != "session" {
		t.Errorf("method[1].Package = %q, want session", methods[1].Package)
	}
}
