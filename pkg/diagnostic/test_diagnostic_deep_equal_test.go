//ff:func feature=orchestrator type=test control=sequence
//ff:what Diagnostic DeepEqual — 동일 값의 Diagnostic 이 reflect.DeepEqual 로 일치
package diagnostic_test

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestDiagnostic_DeepEqual verifies that two Diagnostics with the same values compare equal via reflect.DeepEqual.
func TestDiagnostic_DeepEqual(t *testing.T) {
	a := diagnostic.Diagnostic{
		File:    "a.yaml",
		Line:    1,
		Phase:   diagnostic.PhaseParse,
		Level:   diagnostic.LevelWarning,
		Message: "m",
		Advice:  "adv",
	}
	b := diagnostic.Diagnostic{
		File:    "a.yaml",
		Line:    1,
		Phase:   diagnostic.PhaseParse,
		Level:   diagnostic.LevelWarning,
		Message: "m",
		Advice:  "adv",
	}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("DeepEqual: want equal, got diff\n  a=%+v\n  b=%+v", a, b)
	}

	b.Line = 2
	if reflect.DeepEqual(a, b) {
		t.Errorf("DeepEqual: want diff after Line change, still equal")
	}
}
