//ff:func feature=gen-ir type=test control=sequence
//ff:what convertPublish/resolveExposeInternal/isCountResultType/ddlTableSingularIR/DDLTableSingularIR/findDDLTable
package ir

import (
	"testing"
)

func TestDDLTableSingularIR(t *testing.T) {
	// exported and unexported delegate to caseconv.TableSingular
	if got := DDLTableSingularIR("users"); got != ddlTableSingularIR("users") {
		t.Errorf("exported/unexported mismatch")
	}
	if DDLTableSingularIR("users") == "" {
		t.Errorf("expected non-empty singular")
	}
}
