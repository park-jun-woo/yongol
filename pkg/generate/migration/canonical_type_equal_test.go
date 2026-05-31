//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCanonicalTypeEqual — 모든 필드가 같을 때만 Equal=true
package migration

import (
	"testing"
)

func TestCanonicalTypeEqual(t *testing.T) {
	base := CanonicalType{Base: "VARCHAR", Length: 10, Precision: 0, Scale: 0, Array: false}
	cases := []struct {
		name  string
		other CanonicalType
		want  bool
	}{
		{"identical", base, true},
		{"diff base", CanonicalType{Base: "TEXT", Length: 10}, false},
		{"diff length", CanonicalType{Base: "VARCHAR", Length: 20}, false},
		{"diff precision", CanonicalType{Base: "VARCHAR", Length: 10, Precision: 2}, false},
		{"diff scale", CanonicalType{Base: "VARCHAR", Length: 10, Scale: 2}, false},
		{"diff array", CanonicalType{Base: "VARCHAR", Length: 10, Array: true}, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := base.Equal(c.other); got != c.want {
				t.Errorf("Equal() = %v, want %v", got, c.want)
			}
		})
	}
}
