//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestCatalogHelpers — unit tests for the pure rule catalog helper functions
package catalog

import (
	"testing"
)

func TestIsRuleTableHeader(t *testing.T) {
	if !isRuleTableHeader("| Rule ID | Level | Description | Source |") {
		t.Error("expected canonical header recognised")
	}
	// Case-insensitive.
	if !isRuleTableHeader("| rule id | LEVEL | description | source |") {
		t.Error("expected case-insensitive match")
	}
	// Wrong columns.
	if isRuleTableHeader("| ID | Name |") {
		t.Error("non-header should be rejected")
	}
	// Too few cells.
	if isRuleTableHeader("| Rule ID | Level |") {
		t.Error("too few cells should be rejected")
	}
}
