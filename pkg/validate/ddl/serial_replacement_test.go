//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestSerialReplacement(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"bigserial", "BIGINT GENERATED ALWAYS AS IDENTITY"},
		{"serial", "INTEGER GENERATED ALWAYS AS IDENTITY"},
		{"smallserial", "SMALLINT GENERATED ALWAYS AS IDENTITY"},
		{"unknown", "BIGINT GENERATED ALWAYS AS IDENTITY"},
		{"", "BIGINT GENERATED ALWAYS AS IDENTITY"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := serialReplacement(tt.in); got != tt.want {
				t.Errorf("serialReplacement(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
