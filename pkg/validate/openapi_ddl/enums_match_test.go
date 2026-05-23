//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what enumsMatch — 길이 불일치/순서 무관 일치/값 불일치 검증

package openapi_ddl

import "testing"

func TestEnumsMatch(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want bool
	}{
		{"both nil", nil, nil, true},
		{"both empty", []string{}, []string{}, true},
		{"same elements same order", []string{"a", "b"}, []string{"a", "b"}, true},
		{"same elements different order", []string{"b", "a"}, []string{"a", "b"}, true},
		{"different lengths", []string{"a"}, []string{"a", "b"}, false},
		{"same length different values", []string{"a", "b"}, []string{"a", "c"}, false},
		{"single element match", []string{"x"}, []string{"x"}, true},
		{"single element mismatch", []string{"x"}, []string{"y"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := enumsMatch(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("enumsMatch(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
