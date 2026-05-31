//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestIsSentinelAnnotationLine(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"-- @sentinel", true},
		{"--  @sentinel", true},
		{"--@sentinel", true},
		{"-- @other", false},
		{"id BIGINT", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := isSentinelAnnotationLine(tt.line); got != tt.want {
				t.Errorf("isSentinelAnnotationLine(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}
