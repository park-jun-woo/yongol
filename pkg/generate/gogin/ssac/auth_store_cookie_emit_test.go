//ff:func feature=gen-gogin type=test control=sequence
//ff:what authStoreCookieEmit 단위 테스트 (RefreshRotate Set-Cookie / Logout clear / subscribe no-op)
package ssac

import (
	"strings"
	"testing"
)

func TestAuthStoreCookieEmit(t *testing.T) {
	t.Run("subscribe emits nothing", func(t *testing.T) {
		g := &methodGen{IsSubscribe: true}
		if got := g.authStoreCookieEmit("RefreshRotate", "tok"); got != nil {
			t.Errorf("subscribe should emit nil, got %v", got)
		}
	})
	t.Run("RefreshRotate with blank var → nil", func(t *testing.T) {
		g := &methodGen{}
		if got := g.authStoreCookieEmit("RefreshRotate", "_"); got != nil {
			t.Errorf("blank var should emit nil, got %v", got)
		}
	})
	t.Run("RefreshRotate emits SetAuthCookies", func(t *testing.T) {
		g := &methodGen{}
		joined := strings.Join(g.authStoreCookieEmit("RefreshRotate", "rotated"), "\n")
		if !strings.Contains(joined, "auth.SetAuthCookies(ginCtx, rotated.AccessToken, rotated.RefreshToken)") {
			t.Errorf("missing SetAuthCookies:\n%s", joined)
		}
	})
	t.Run("Logout clears cookies", func(t *testing.T) {
		g := &methodGen{}
		joined := strings.Join(g.authStoreCookieEmit("Logout", "_"), "\n")
		if !strings.Contains(joined, "auth.ClearAuthCookies(ginCtx)") {
			t.Errorf("missing ClearAuthCookies:\n%s", joined)
		}
	})
	t.Run("unknown call → nil", func(t *testing.T) {
		g := &methodGen{}
		if got := g.authStoreCookieEmit("Other", "x"); got != nil {
			t.Errorf("unknown call should emit nil, got %v", got)
		}
	})
}
