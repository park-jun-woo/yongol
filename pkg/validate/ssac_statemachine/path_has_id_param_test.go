//ff:func feature=validate type=test control=selection topic=states
//ff:what TestPathHasIDParam — pathHasIDParam {id} 파라미터 포함 여부 분기 검증

package ssac_statemachine

import "testing"

func TestPathHasIDParam(t *testing.T) {
	cases := []struct {
		path string
		want bool
	}{
		{"/orders/{id}", true},
		{"/orders/{ID}", true},
		{"/orders/{Id}", true},
		{"/orders/{slug}", false},
		{"/orders", false},
		{"/orders/{id", false},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			if got := pathHasIDParam(tc.path); got != tc.want {
				t.Errorf("pathHasIDParam(%q) = %v, want %v", tc.path, got, tc.want)
			}
		})
	}
}
