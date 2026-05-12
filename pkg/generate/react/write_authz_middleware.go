//ff:func feature=gen-react type=generator control=sequence
//ff:what writeAuthzMiddleware — openapi-fetch 클라이언트에 JWT 인증 미들웨어 코드를 방출한다
package react

import "strings"

// writeAuthzMiddleware appends the JWT auth middleware snippet to b.
// It registers onRequest (attach Bearer token) and onResponse (handle 401).
func writeAuthzMiddleware(b *strings.Builder) {
	b.WriteString("\nclient.use({\n")
	b.WriteString("  async onRequest({ request }) {\n")
	b.WriteString("    const token = localStorage.getItem('access_token')\n")
	b.WriteString("    if (token) {\n")
	b.WriteString("      request.headers.set('Authorization', `Bearer ${token}`)\n")
	b.WriteString("    }\n")
	b.WriteString("    return request\n")
	b.WriteString("  },\n")
	b.WriteString("  async onResponse({ response }) {\n")
	b.WriteString("    if (response.status === 401) {\n")
	b.WriteString("      localStorage.removeItem('access_token')\n")
	b.WriteString("      localStorage.removeItem('refresh_token')\n")
	b.WriteString("      window.location.href = '/login'\n")
	b.WriteString("    }\n")
	b.WriteString("    return response\n")
	b.WriteString("  },\n")
	b.WriteString("})\n")
}
