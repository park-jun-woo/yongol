//ff:func feature=manifest type=parser control=sequence
//ff:what BuildLineIndex 가 첫 2xx response 필드 줄 번호를 올바르게 색인하는지 검증

package openapi

import "testing"

func TestBuildLineIndex_ResponseFields(t *testing.T) {
	path := writeFixture(t)
	idx, _ := BuildLineIndex(path)
	if got, want := idx.ResponseFieldLine("Login", "access_token"), 35; got != want {
		t.Errorf("ResponseFieldLine(Login,access_token) = %d, want %d", got, want)
	}
}
