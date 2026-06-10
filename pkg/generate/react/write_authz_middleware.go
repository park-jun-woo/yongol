//ff:func feature=gen-react type=generator control=sequence
//ff:what writeAuthzMiddleware — 세션 store 기반 Bearer 주입 미들웨어 방출 (refresh 시 주입 전용)

package react

import "strings"

// writeAuthzMiddleware appends the bearer auth middleware snippet to b.
// onRequest attaches the Bearer token read from the session store
// (src/stores/auth.ts — emitted under the same bearer gate).
//
// The 401 semantics depend on withRefresh (Phase004):
//   - true: injection only. 401 handling lives in the withAuthRetry
//     operation wrapper (writeRefreshFlow) — the openapi-fetch middleware
//     cannot retry the original request from onResponse.
//   - false: the explicit downgrade (frontend.auth.refresh_field
//     undeclared) — onResponse keeps the pre-Phase004 semantics, 401 →
//     clear store + redirect to /login.
func writeAuthzMiddleware(b *strings.Builder, withRefresh bool) {
	b.WriteString("\nclient.use({\n")
	b.WriteString("  async onRequest({ request }) {\n")
	b.WriteString("    const token = useAuthStore.getState().token\n")
	b.WriteString("    if (token) {\n")
	b.WriteString("      request.headers.set('Authorization', `Bearer ${token}`)\n")
	b.WriteString("    }\n")
	b.WriteString("    return request\n")
	b.WriteString("  },\n")
	if !withRefresh {
		b.WriteString("  async onResponse({ response }) {\n")
		b.WriteString("    if (response.status === 401) {\n")
		b.WriteString("      useAuthStore.getState().clear()\n")
		b.WriteString("      window.location.href = '/login'\n")
		b.WriteString("    }\n")
		b.WriteString("    return response\n")
		b.WriteString("  },\n")
	}
	b.WriteString("})\n")
}
