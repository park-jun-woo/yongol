//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"reflect"
	"testing"
)

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want []string
	}{
		{"single string", "a", []string{"a"}},
		{"empty string", "", []string{""}},
		{"slice of strings", []interface{}{"a", "b"}, []string{"a", "b"}},
		{"slice with non-strings", []interface{}{"a", 1, "b"}, []string{"a", "b"}},
		{"slice all non-strings", []interface{}{1, 2}, nil},
		{"unrecognised type int", 42, nil},
		{"unrecognised type nil", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toStringSlice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringSlice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
