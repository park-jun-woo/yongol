//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"reflect"
	"testing"
)

func TestExtractStringsFromSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []interface{}
		want []string
	}{
		{"all strings", []interface{}{"x", "y"}, []string{"x", "y"}},
		{"mixed", []interface{}{"x", 3, true, "y"}, []string{"x", "y"}},
		{"none", []interface{}{1, 2, 3}, nil},
		{"empty", []interface{}{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractStringsFromSlice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
