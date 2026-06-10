//ff:func feature=gen-react type=test control=sequence
//ff:what writeCookie401Redirect — cookie 모드 401 → /login 수렴 미들웨어 방출 검증

package react

import (
	"strings"
	"testing"
)

func TestWriteCookie401Redirect(t *testing.T) {
	var b strings.Builder
	writeCookie401Redirect(&b)
	content := b.String()

	assertContains(t, content, "async onResponse({ response })")
	assertContains(t, content, "response.status === 401")
	assertContains(t, content, "window.location.pathname !== '/login'")
	assertContains(t, content, "window.location.href = '/login'")
}
