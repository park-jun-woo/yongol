//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — nil OpenAPIDoc 시 진단 0

package openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case1_NilDoc(t *testing.T) {
	diags := o06RequiredPropertyConsistency(&yongol.Fullstack{})
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
