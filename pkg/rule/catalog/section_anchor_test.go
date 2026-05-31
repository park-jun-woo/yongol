//ff:func feature=rule type=test control=iteration dimension=1 topic=catalog
//ff:what TestCatalogHelpers — unit tests for the pure rule catalog helper functions
package catalog

import (
	"testing"
)

func TestSectionAnchor(t *testing.T) {
	tests := []struct{ in, want string }{
		{"D-Series Rules", "d-series-rules"},
		{"XQS.20 Return Type", "xqs-20-return-type"},
		{"Trailing  ", "trailing"},
		{"a/b c", "a-b-c"},
		{"!!!", ""}, // all punctuation stripped
	}
	for _, tt := range tests {
		if got := sectionAnchor(tt.in); got != tt.want {
			t.Errorf("sectionAnchor(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
