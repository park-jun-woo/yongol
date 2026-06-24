//ff:func feature=gen-gogin type=test control=sequence
//ff:what authMwFileName — 단일 사이트 고정명 vs 도메인 접미 파일명 분기 검증

package auth

import "testing"

func TestAuthMwFileName(t *testing.T) {
	if got := authMwFileName("bearerauth", ""); got != "bearerauth.go" {
		t.Errorf("single-site → bearerauth.go, got %q", got)
	}
	if got := authMwFileName("bearerauth", "admin"); got != "bearerauth_admin.go" {
		t.Errorf("domain → bearerauth_admin.go, got %q", got)
	}
	if got := authMwFileName("cookieauth", "public"); got != "cookieauth_public.go" {
		t.Errorf("domain → cookieauth_public.go, got %q", got)
	}
}
