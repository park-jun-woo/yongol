//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what Testtruthy — truthy() Go 기본 타입별 zero/non-zero 판정 테이블 테스트

package ssacmeta

import "testing"

func TestTruthy(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want bool
	}{
		{"nil", nil, false},
		{"bool-true", true, true},
		{"bool-false", false, false},
		{"string-nonempty", "x", true},
		{"string-empty", "", false},
		{"int-nonzero", 3, true},
		{"int-zero", 0, false},
		{"int64-nonzero", int64(7), true},
		{"int64-zero", int64(0), false},
		{"float64-nonzero", 1.5, true},
		{"float64-zero", 0.0, false},
		{"default-map", map[string]any{}, true},
		{"default-slice", []int{}, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := truthy(c.in); got != c.want {
				t.Errorf("truthy(%#v) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}
