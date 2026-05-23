//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what diffStringSets — nil/empty/부분/완전 차집합 정렬 검증

package openapi

import (
	"reflect"
	"testing"
)

func TestDiffStringSets(t *testing.T) {
	tests := []struct {
		name string
		a    map[string]bool
		b    map[string]bool
		want []string
	}{
		{
			name: "both nil returns empty",
			a:    nil,
			b:    nil,
			want: nil,
		},
		{
			name: "a empty returns empty",
			a:    map[string]bool{},
			b:    map[string]bool{"x": true},
			want: nil,
		},
		{
			name: "b empty returns all of a sorted",
			a:    map[string]bool{"z": true, "a": true, "m": true},
			b:    map[string]bool{},
			want: []string{"a", "m", "z"},
		},
		{
			name: "no difference returns empty",
			a:    map[string]bool{"x": true, "y": true},
			b:    map[string]bool{"x": true, "y": true},
			want: nil,
		},
		{
			name: "partial difference",
			a:    map[string]bool{"x": true, "y": true, "z": true},
			b:    map[string]bool{"y": true},
			want: []string{"x", "z"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffStringSets(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
