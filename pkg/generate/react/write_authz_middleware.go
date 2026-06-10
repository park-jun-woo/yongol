//ff:func feature=gen-react type=generator control=sequence
//ff:what writeAuthzMiddleware — openapi-fetch 클라이언트에 세션 store 기반 Bearer 인증 미들웨어 방출
package react

import "strings"

// writeAuthzMiddleware appends the bearer auth middleware snippet to b.
// It registers onRequest (attach Bearer token read from the session store)
// and onResponse (401 → clear store + redirect to /login). Emitted only in
// bearer mode — the store it reads (src/stores/auth.ts) is emitted under
// the same gate. The 401 semantics (clear + /login) are unchanged in
// Phase003; the single-flight refresh swap is Phase004.
func writeAuthzMiddleware(b *strings.Builder) {
	b.WriteString("\nclient.use({\n")
	b.WriteString("  async onRequest({ request }) {\n")
	b.WriteString("    const token = useAuthStore.getState().token\n")
	b.WriteString("    if (token) {\n")
	b.WriteString("      request.headers.set('Authorization', `Bearer ${token}`)\n")
	b.WriteString("    }\n")
	b.WriteString("    return request\n")
	b.WriteString("  },\n")
	b.WriteString("  async onResponse({ response }) {\n")
	b.WriteString("    if (response.status === 401) {\n")
	b.WriteString("      useAuthStore.getState().clear()\n")
	b.WriteString("      window.location.href = '/login'\n")
	b.WriteString("    }\n")
	b.WriteString("    return response\n")
	b.WriteString("  },\n")
	b.WriteString("})\n")
}
