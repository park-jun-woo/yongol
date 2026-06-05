//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO12ItemDiags — no-front인데 소비된 operation에만 XMO-12 WARNING 생성, 미소비/non-no-front/빈ID/nil verb 스킵 검증
package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestXMO12ItemDiags(t *testing.T) {
	item := &openapi3.PathItem{
		Get:    &openapi3.Operation{OperationID: "NoFrontConsumed", Tags: []string{noFrontTag}},
		Post:   &openapi3.Operation{OperationID: "NoFrontUnconsumed", Tags: []string{noFrontTag}},
		Put:    &openapi3.Operation{OperationID: "PlainConsumed"},
		Delete: &openapi3.Operation{OperationID: "", Tags: []string{noFrontTag}},
		// Patch is nil (undefined verb) → skipped.
	}

	consumed := map[string]struct{}{
		"NoFrontConsumed": {},
		"PlainConsumed":   {},
	}

	diags := xmo12ItemDiags(item, consumed)

	// Only NoFrontConsumed triggers an XMO-12 WARNING.
	if got := countDiag(diags, "[XMO-12]"); got != 1 {
		t.Fatalf("expected 1 XMO-12 diag, got %d: %+v", got, diags)
	}
	d := diags[0]
	if d.Level != diagnostic.LevelWarning {
		t.Errorf("level = %q, want WARNING", d.Level)
	}
	if d.File != "openapi.yaml" || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("unexpected File/Phase: %q/%q", d.File, d.Phase)
	}
	if !strings.Contains(d.Message, "NoFrontConsumed") {
		t.Errorf("message should name NoFrontConsumed, got %q", d.Message)
	}

	// no-front but unconsumed, and plain consumed → no XMO-12 diags.
	if strings.Contains(d.Message, "NoFrontUnconsumed") || strings.Contains(d.Message, "PlainConsumed") {
		t.Errorf("unexpected op in message: %q", d.Message)
	}

	// Nothing consumed → no diagnostics.
	if got := xmo12ItemDiags(item, map[string]struct{}{}); len(got) != 0 {
		t.Errorf("none consumed: expected 0 diags, got %d", len(got))
	}
}
