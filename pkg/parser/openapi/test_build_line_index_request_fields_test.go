//ff:func feature=manifest type=parser control=sequence
//ff:what BuildLineIndex 가 requestBody 필드 줄 번호를 올바르게 색인하는지 검증

package openapi

import "testing"

func TestBuildLineIndex_RequestFields(t *testing.T) {
	path := writeFixture(t)
	idx, _ := BuildLineIndex(path)
	// "email: { type: string, format: email, maxLength: 255 }" 줄
	if got, want := idx.RequestFieldLine("Login", "email"), 25; got != want {
		t.Errorf("RequestFieldLine(Login,email) = %d, want %d", got, want)
	}
	// "password: { type: string, minLength: 8 }" 줄
	if got, want := idx.RequestFieldLine("Login", "password"), 26; got != want {
		t.Errorf("RequestFieldLine(Login,password) = %d, want %d", got, want)
	}
}
