//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what TestGroundHelpers — unit tests for the pure ground helper functions
package ground

import (
	"testing"
)

func TestIsResponseCodeRelevant(t *testing.T) {
	tests := []struct {
		code string
		want bool
	}{
		{"200", true},
		{"404", true},
		{"500", true},
		{"100", false},
		{"301", false},
	}
	for _, tt := range tests {
		if got := isResponseCodeRelevant(tt.code); got != tt.want {
			t.Errorf("isResponseCodeRelevant(%q) = %v, want %v", tt.code, got, tt.want)
		}
	}
}
