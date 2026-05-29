//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what TestKebabToCamel — kebab-case to camelCase conversion table test

package stml

import (
	"testing"
)

func TestKebabToCamel(t *testing.T) {
	tests := []struct{ in, want string }{
		{"project-id", "projectId"},
		{"ReservationID", "ReservationID"},
		{"room-id", "roomId"},
		{"a-b-c", "aBC"},
	}
	for _, tt := range tests {
		got := kebabToCamel(tt.in)
		if got != tt.want {
			t.Errorf("kebabToCamel(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
