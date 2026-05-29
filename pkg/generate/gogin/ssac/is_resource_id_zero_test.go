//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isResourceIDZero 단위 테스트

package ssac

import "testing"

func TestIsResourceIDZero(t *testing.T) {
	cases := map[string]bool{
		"":            true,
		"   ":         true,
		"0":           true,
		`""`:          true,
		"''":          true,
		"nil":         true,
		"null":        true,
		"NULL":        true,
		" 0 ":         true,
		"resource.ID": false,
		"1":           false,
	}
	for in, want := range cases {
		if got := isResourceIDZero(in); got != want {
			t.Errorf("isResourceIDZero(%q) = %v, want %v", in, got, want)
		}
	}
}
