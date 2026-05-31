//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestRiskyCast — 같은 Base 축소 또는 카테고리 교차 변환은 risky
package migration

import (
	"testing"
)

func TestRiskyCast(t *testing.T) {
	cases := []struct {
		name string
		from CanonicalType
		to   CanonicalType
		want bool
	}{
		{"varchar shrink", CanonicalType{Base: "VARCHAR", Length: 255}, CanonicalType{Base: "VARCHAR", Length: 50}, true},
		{"varchar grow", CanonicalType{Base: "VARCHAR", Length: 50}, CanonicalType{Base: "VARCHAR", Length: 255}, false},
		{"int to text", CanonicalType{Base: "INTEGER"}, CanonicalType{Base: "TEXT"}, true},
		{"int widen", CanonicalType{Base: "INTEGER"}, CanonicalType{Base: "BIGINT"}, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := riskyCast(c.from, c.to); got != c.want {
				t.Errorf("riskyCast = %v, want %v", got, c.want)
			}
		})
	}
}
