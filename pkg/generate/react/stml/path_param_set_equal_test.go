//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what pathParamSetEqual — 동일 이름집합/길이불일치/키불일치/빈집합 비교 검증

package stml

import "testing"

func TestPathParamSetEqual(t *testing.T) {
	tests := []struct {
		name string
		a, b map[string]string
		want bool
	}{
		{
			name: "same names ignoring types",
			a:    map[string]string{"id": "integer", "slug": "string"},
			b:    map[string]string{"id": "string", "slug": "integer"},
			want: true,
		},
		{
			name: "different lengths",
			a:    map[string]string{"id": "integer"},
			b:    map[string]string{"id": "integer", "slug": "string"},
			want: false,
		},
		{
			name: "same length different keys",
			a:    map[string]string{"id": "integer"},
			b:    map[string]string{"slug": "string"},
			want: false,
		},
		{
			name: "both empty",
			a:    map[string]string{},
			b:    map[string]string{},
			want: false,
		},
		{
			name: "a empty b non-empty",
			a:    nil,
			b:    map[string]string{"id": "integer"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathParamSetEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("pathParamSetEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
