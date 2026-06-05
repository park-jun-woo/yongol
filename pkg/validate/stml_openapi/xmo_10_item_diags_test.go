//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10ItemDiags — 미소비·non-no-front operation에만 XMO-10 ERROR 생성, 소비/no-front/빈ID/nil verb 스킵 검증
package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestXMO10ItemDiags(t *testing.T) {
	item := &openapi3.PathItem{
		Get:    &openapi3.Operation{OperationID: "ConsumedOp"},
		Post:   &openapi3.Operation{OperationID: "UnconsumedOp"},
		Put:    &openapi3.Operation{OperationID: "NoFrontOp", Tags: []string{noFrontTag}},
		Delete: &openapi3.Operation{OperationID: ""},
		// Patch is nil (undefined verb) → skipped.
	}

	consumed := map[string]struct{}{"ConsumedOp": {}}

	diags := xmo10ItemDiags(item, consumed)

	// Only UnconsumedOp should trigger an XMO-10 ERROR.
	if got := countDiag(diags, "[XMO-10]"); got != 1 {
		t.Fatalf("expected 1 XMO-10 diag, got %d: %+v", got, diags)
	}
	d := diags[0]
	if d.Level != diagnostic.LevelError {
		t.Errorf("level = %q, want ERROR", d.Level)
	}
	if d.File != "openapi.yaml" || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("unexpected File/Phase: %q/%q", d.File, d.Phase)
	}
	if !strings.Contains(d.Message, "UnconsumedOp") {
		t.Errorf("message should name UnconsumedOp, got %q", d.Message)
	}
	if strings.Contains(d.Message, "ConsumedOp") || strings.Contains(d.Message, "NoFrontOp") {
		t.Errorf("consumed/no-front ops must not appear, got %q", d.Message)
	}

	// All consumed → no diagnostics.
	allConsumed := map[string]struct{}{"ConsumedOp": {}, "UnconsumedOp": {}}
	if got := xmo10ItemDiags(item, allConsumed); len(got) != 0 {
		t.Errorf("all consumed: expected 0 diags, got %d", len(got))
	}
}
