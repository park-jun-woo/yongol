//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"strings"
	"testing"
)

func TestByNameWrapAuthClaimsFields_ZeroCov(t *testing.T) {
	if got := wrapAuthClaimsFields("auth", "IssueToken", "UserID: id"); !strings.Contains(got, "model.UserClaim{") {
		t.Errorf("issue token wrap = %q", got)
	}
	if got := wrapAuthClaimsFields("auth", "RefreshToken", "UserID: id"); !strings.Contains(got, "model.UserClaim{") {
		t.Errorf("refresh token wrap = %q", got)
	}
	// non-auth passthrough.
	if got := wrapAuthClaimsFields("mail", "Send", "To: x"); got != "To: x" {
		t.Errorf("passthrough = %q", got)
	}
}
