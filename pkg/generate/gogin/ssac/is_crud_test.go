//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isCRUD 단위 테스트 (get/post/put/delete 만 CRUD)

package ssac

import "testing"

func TestIsCRUD(t *testing.T) {
	cases := map[string]bool{
		"get":       true,
		"post":      true,
		"put":       true,
		"delete":    true,
		"call":      false,
		"eval":      false,
		"publish":   false,
		"subscribe": false,
		"":          false,
	}
	for in, want := range cases {
		if got := isCRUD(in); got != want {
			t.Errorf("isCRUD(%q) = %v, want %v", in, got, want)
		}
	}
}
