//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"
)

func TestXsp29MissingSSaCDiag(t *testing.T) {
	pair := [2]string{"read", "doc"}
	pairLoc := map[[2]string]PairLocation{pair: {File: "p.rego", Line: 8}}

	diag, ok := xsp29MissingSSaCDiag(pair, map[[2]string]bool{}, pairLoc)
	if !ok {
		t.Fatal("expected diagnostic for missing SSaC @auth")
	}
	if diag.Line != 8 || diag.File != "p.rego" {
		t.Errorf("diag = %+v", diag)
	}
	if _, ok := xsp29MissingSSaCDiag(pair, map[[2]string]bool{pair: true}, pairLoc); ok {
		t.Error("present pair should not produce a diagnostic")
	}
}
