//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"strings"
	"testing"
)

func TestEmitIfClass(t *testing.T) {
	diags := emitIfClass("page.stml", "div", "card", "bg-red")
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[TM-10]") || !strings.Contains(diags[0].Message, "bg-red") {
		t.Errorf("unexpected message: %q", diags[0].Message)
	}
	// empty className → no diagnostic.
	if got := emitIfClass("page.stml", "div", "card", ""); got != nil {
		t.Errorf("empty class should yield nil, got %v", got)
	}
}
