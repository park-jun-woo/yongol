//ff:func feature=gen-react type=generator control=sequence
//ff:what writeCookie401Redirect — cookie 모드 401 → /login 수렴 미들웨어 방출

package react

import "strings"

// writeCookie401Redirect appends the cookie-mode 401 convergence middleware
// to api.ts. The route guard passes optimistically in cookie mode (the
// httpOnly session is unreadable client-side — Phase005 ProtectedRoute), so
// an absent/expired session surfaces as a 401 on the first protected fetch;
// this middleware converges it to /login. Requests issued from /login itself
// are exempt so a failed login's 401 keeps its inline data-on-error message
// instead of reloading the page.
func writeCookie401Redirect(b *strings.Builder) {
	b.WriteString(`
// Cookie-mode 401 convergence: the session is an httpOnly cookie the guard
// cannot inspect, so an unauthenticated protected fetch answers 401 here.
client.use({
  async onResponse({ response }) {
    if (response.status === 401 && window.location.pathname !== '/login') {
      window.location.href = '/login'
    }
    return response
  },
})
`)
}
