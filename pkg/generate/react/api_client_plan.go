//ff:type feature=gen-react type=model
//ff:what apiClientPlan — api.ts 모드 분기 계획 (bearer 주입·refresh / cookie credentials·CSRF)

package react

// apiClientPlan drives the writeAPIClient mode branches (plans/stml/auth-flow
// Phase004). At most one of bearer/cookie is true: bearer when the prepared
// auth mode is "bearer", cookie when it is "cookie" or "hybrid" (the browser
// session rides httpOnly cookies in both). Both false means backend.auth is
// absent and a plain client is emitted.
type apiClientPlan struct {
	// bearer gates the session-store import and the Bearer-injection
	// middleware. The 401 semantics depend on refresh below.
	bearer bool
	// cookie gates `credentials: 'include'` on createClient so the
	// httpOnly session cookies travel with every request.
	cookie bool
	// csrf gates the double-submit CSRF request middleware (cookie path
	// only). False when backend.auth.csrf.enabled is explicitly false —
	// mirrors ir.csrfIsActive so the frontend never attaches a header the
	// backend middleware would not verify.
	csrf bool
	// csrfCookieName is the JS-readable cookie the backend CSRF
	// middleware sets (backend.auth.csrf.cookie_name, default XSRF-TOKEN).
	csrfCookieName string
	// csrfHeaderName is the header the backend verifies on state-changing
	// requests (backend.auth.csrf.header_name, default X-XSRF-TOKEN).
	csrfHeaderName string
	// refresh is the bearer 401→refresh→retry plan. nil in bearer mode is
	// the explicit downgrade (frontend.auth.refresh_field undeclared or no
	// usable refresh op): Bearer injection only, 401 clears the store and
	// redirects to /login.
	refresh *refreshPlan
}
