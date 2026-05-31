//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestCatalogHelpers — unit tests for the pure rule catalog helper functions
package catalog

import (
	"testing"
)

func TestIsTableSeparator(t *testing.T) {
	if !isTableSeparator("|---|---|---|---|") {
		t.Error("expected separator recognised")
	}
	if !isTableSeparator("| :--- | ---: | :---: |") {
		t.Error("alignment colons should be accepted")
	}
	// Empty cell → not a separator.
	if isTableSeparator("|---||---|") {
		t.Error("empty cell should reject")
	}
	// Non-dash content → not a separator.
	if isTableSeparator("| abc | def |") {
		t.Error("text content should reject")
	}
}
