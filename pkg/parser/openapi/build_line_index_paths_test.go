//ff:func feature=manifest type=parser control=sequence
//ff:what BuildLineIndex 가 paths: 아래 각 path key 줄 번호를 올바르게 색인하는지 검증

package openapi

import "testing"

func TestBuildLineIndex_Paths(t *testing.T) {
	path := writeFixture(t)
	idx, _ := BuildLineIndex(path)
	if got, want := idx.Paths["/login"], 14; got != want {
		t.Errorf("Paths[/login] = %d, want %d", got, want)
	}
	if got, want := idx.Paths["/users/me"], 36; got != want {
		t.Errorf("Paths[/users/me] = %d, want %d", got, want)
	}
}
