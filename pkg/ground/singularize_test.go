//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what TestGroundHelpers — unit tests for the pure ground helper functions
package ground

import (
	"testing"
)

func TestSingularize(t *testing.T) {
	tests := []struct{ in, want string }{
		{"users", "user"},
		{"categories", "category"},  // ies → y
		{"addresses", "address"},    // sses → ss
		{"boxes", "box"},            // xes → x
		{"audit_logs", "audit_log"}, // plain s
		{"address", "address"},      // ends in ss → unchanged
		{"workflow", "workflow"},    // no plural suffix
	}
	for _, tt := range tests {
		if got := singularize(tt.in); got != tt.want {
			t.Errorf("singularize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
