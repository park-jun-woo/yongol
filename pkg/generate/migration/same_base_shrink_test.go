//ff:func feature=migration type=test control=sequence
//ff:what TestSameBaseShrink — 같은 Base VARCHAR/NUMERIC 축소 판정
package migration

import "testing"

func TestSameBaseShrink(t *testing.T) {
	cases := []struct {
		name string
		from CanonicalType
		to   CanonicalType
		want bool
	}{
		{"varchar shrink", CanonicalType{Base: "VARCHAR", Length: 255}, CanonicalType{Base: "VARCHAR", Length: 100}, true},
		{"varchar grow", CanonicalType{Base: "VARCHAR", Length: 100}, CanonicalType{Base: "VARCHAR", Length: 255}, false},
		{"numeric shrink", CanonicalType{Base: "NUMERIC", Precision: 10}, CanonicalType{Base: "NUMERIC", Precision: 5}, true},
		{"numeric grow", CanonicalType{Base: "NUMERIC", Precision: 5}, CanonicalType{Base: "NUMERIC", Precision: 10}, false},
		{"text no length", CanonicalType{Base: "TEXT"}, CanonicalType{Base: "TEXT"}, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := sameBaseShrink(c.from, c.to); got != c.want {
				t.Errorf("sameBaseShrink = %v, want %v", got, c.want)
			}
		})
	}
}
