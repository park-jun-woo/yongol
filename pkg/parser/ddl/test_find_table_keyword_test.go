//ff:func feature=manifest type=test control=sequence
//ff:what findTableKeyword — ["CREATE","TABLE",...] 토큰 리스트에서 TABLE 인덱스 반환

package ddl

import "testing"

func TestFindTableKeyword(t *testing.T) {
	if got := findTableKeyword([]string{"CREATE", "TABLE", "users", "("}); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	if got := findTableKeyword([]string{"CREATE", "INDEX", "idx"}); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}
