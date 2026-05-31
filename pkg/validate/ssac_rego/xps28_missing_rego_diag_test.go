//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"
)

func TestXps28MissingRegoDiag(t *testing.T) {
	pair := [2]string{"delete", "project"}
	pairLoc := map[[2]string]PairLocation{pair: {File: "s.ssac", Line: 12}}

	// Pair absent from rego → diagnostic.
	diag, ok := xps28MissingRegoDiag(pair, map[[2]string]bool{}, pairLoc)
	if !ok {
		t.Fatal("expected diagnostic for missing rego rule")
	}
	if diag.Line != 12 || diag.OperationID != "delete" {
		t.Errorf("diag = %+v", diag)
	}
	// Pair present → no diagnostic.
	if _, ok := xps28MissingRegoDiag(pair, map[[2]string]bool{pair: true}, pairLoc); ok {
		t.Error("present pair should not produce a diagnostic")
	}
}
