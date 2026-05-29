//ff:func feature=migration type=test control=sequence
//ff:what TestPKName — 기본 PK 제약 이름은 <table>_pkey (소문자)
package migration

import "testing"

func TestPKName(t *testing.T) {
	if got := PKName("users"); got != "users_pkey" {
		t.Errorf("PKName(users) = %q, want users_pkey", got)
	}
	if got := PKName("Posts"); got != "posts_pkey" {
		t.Errorf("PKName(Posts) = %q, want posts_pkey", got)
	}
}
