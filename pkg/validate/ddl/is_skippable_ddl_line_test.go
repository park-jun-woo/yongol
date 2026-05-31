//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestIsSkippableDDLLine(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"", true},
		{"-- comment", true},
		{");", true},
		{"CREATE TABLE foo (", true},
		{"INSERT INTO foo VALUES (1)", true},
		{"ON CONFLICT DO NOTHING", true},
		{"VALUES (1, 2)", true},
		{"PRIMARY KEY (id)", true},
		{"UNIQUE (email)", true},
		{"CHECK (x > 0)", true},
		{"FOREIGN KEY (a) REFERENCES b", true},
		{"CONSTRAINT fk_x FOREIGN KEY", true},
		{"id BIGINT", false},
		{"name TEXT NOT NULL", false},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := isSkippableDDLLine(tt.line); got != tt.want {
				t.Errorf("isSkippableDDLLine(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}
