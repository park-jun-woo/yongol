//ff:func feature=validate type=test control=selection topic=ssac-statemachine
//ff:what TestStateTypesCompatible — stateTypesCompatible 대입 가능성 판별 분기 검증

package ssac_statemachine

import "testing"

func TestStateTypesCompatible(t *testing.T) {
	cases := []struct {
		actual   string
		expected string
		want     bool
	}{
		{"string", "string", true},
		{"*string", "string", true},
		{"nil", "int", true},
		{"int", "nil", true},
		{"int", "string", false},
	}
	for _, tc := range cases {
		t.Run(tc.actual+"_"+tc.expected, func(t *testing.T) {
			if got := stateTypesCompatible(tc.actual, tc.expected); got != tc.want {
				t.Errorf("stateTypesCompatible(%q,%q) = %v, want %v", tc.actual, tc.expected, got, tc.want)
			}
		})
	}
}
