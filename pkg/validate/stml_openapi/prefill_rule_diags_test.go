//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what prefillRuleDiags — TM-54/55/56 묶음이 한 페이지에서 함께 실행되는지 검증

package stml_openapi

import (
	"testing"
)

func TestPrefillRuleDiags(t *testing.T) {
	opMap := buildOperationMethodMap(tm54Doc())

	// Edit page with a GET-by-id fetch but no data-prefill: TM-55 fires (the
	// no-prefill edit-form advisory); TM-54 and TM-56 stay silent (no prefill,
	// PUT not PATCH).
	page := tm54Page(t, ``, `<input data-field="sheet_name" />`)
	diags := prefillRuleDiags(page, opMap)
	if countDiag(diags, "[TM-55]") != 1 {
		t.Fatalf("expected TM-55 to fire, got %+v", diags)
	}
	if countDiag(diags, "[TM-54]") != 0 || countDiag(diags, "[TM-56]") != 0 {
		t.Errorf("TM-54/56 should be silent here, got %+v", diags)
	}

	// Fully wired prefill page: all three silent.
	good := tm54Page(t, ` data-prefill="GetRule"`, `<input data-field="sheet_name" /><input data-field="start_row" />`)
	if d := prefillRuleDiags(good, opMap); len(d) != 0 {
		t.Errorf("fully wired prefill page should be silent, got %+v", d)
	}
}
