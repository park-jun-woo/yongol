//ff:func feature=gen-react type=test control=sequence
//ff:what writeCSRFMiddleware — 쿠키→헤더 미러링 미들웨어 방출(정규식 메타 이스케이프·메서드 게이트 포함) 검증

package react

import (
	"strings"
	"testing"
)

func TestWriteCSRFMiddleware(t *testing.T) {
	var b strings.Builder
	// cookie name carries a regex-special '.' to exercise regexp.QuoteMeta
	writeCSRFMiddleware(&b, "X.SRF", "X-My-Header")
	out := b.String()

	// cookie read uses the QuoteMeta-escaped name in the match regex
	assertContains(t, out, `document.cookie.match(/(?:^|;\s*)X\.SRF=([^;]*)/)`)
	// header mirrored on state-changing requests
	assertContains(t, out, `request.headers.set('X-My-Header', token)`)
	// safe-method gate (complement of backend isSafeMethod)
	assertContains(t, out, `['GET', 'HEAD', 'OPTIONS'].includes(request.method)`)
	// middleware registered on the client
	assertContains(t, out, "client.use({")
}
