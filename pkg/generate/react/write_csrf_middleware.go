//ff:func feature=gen-react type=generator control=sequence
//ff:what writeCSRFMiddleware — double-submit 쿠키 CSRF 헤더 미러링 미들웨어 방출 (cookie 모드)

package react

import (
	"fmt"
	"regexp"
	"strings"
)

// writeCSRFMiddleware appends the cookie-mode CSRF middleware to api.ts.
// It mirrors the backend's double-submit defense
// (pkg/generate/gogin/middleware/csrf_source.go): the server sets the
// JS-readable cookieName cookie on safe requests and verifies headerName on
// state-changing ones, so the client reads the cookie at request time and
// copies it into the header. No memory copy is kept — the cookie is the
// storage. The method gate ("everything but GET/HEAD/OPTIONS") is the exact
// complement of the backend's isSafeMethod, PATCH included.
func writeCSRFMiddleware(b *strings.Builder, cookieName, headerName string) {
	b.WriteString(fmt.Sprintf(`
// Double-submit cookie CSRF: the backend sets the JS-readable %s
// cookie and verifies the %s header on state-changing requests
// (internal/middleware/csrf.go). Mirror cookie -> header per request.
function csrfToken(): string | null {
  const m = document.cookie.match(/(?:^|;\s*)%s=([^;]*)/)
  return m ? decodeURIComponent(m[1]) : null
}

client.use({
  async onRequest({ request }) {
    if (!['GET', 'HEAD', 'OPTIONS'].includes(request.method)) {
      const token = csrfToken()
      if (token) {
        request.headers.set('%s', token)
      }
    }
    return request
  },
})
`, cookieName, headerName, regexp.QuoteMeta(cookieName), headerName))
}
