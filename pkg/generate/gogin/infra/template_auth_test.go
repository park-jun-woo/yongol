//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestAuthWrapperMethodHeader — auth wrapper 메서드 헤더 생성 검증

package infra

import (
	"strings"
	"testing"
)

func TestAuthWrapperMethodHeader(t *testing.T) {
	got := authWrapperMethodHeader("Save", "Save — refresh token 저장")
	if !strings.Contains(got, "//ff:func feature=infra type=accessor control=sequence topic=auth-refresh") {
		t.Errorf("missing ff:func header:\n%s", got)
	}
	if !strings.Contains(got, "//ff:what Save — refresh token 저장") {
		t.Errorf("missing ff:what with passed text:\n%s", got)
	}
}
