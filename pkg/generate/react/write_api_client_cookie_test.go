//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient cookie 모드 — credentials include + double-submit CSRF 헤더 미러링 방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_Cookie_CSRF(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{
			Get:  buildTokenOp("ListItems", []string{"items"}, nil),
			Post: buildTokenOp("CreateItem", []string{"id"}, []string{"name"}),
		}),
	)}
	plan := apiClientPlan{cookie: true, csrf: true, csrfCookieName: "XSRF-TOKEN", csrfHeaderName: "X-XSRF-TOKEN"}
	if err := writeAPIClient(dir, doc, plan); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// httpOnly session cookies ride along on every request.
	assertContains(t, content, "createClient<paths>({ baseUrl: '', credentials: 'include' })")

	// Double-submit CSRF: read the JS-readable cookie, mirror it into the
	// header on state-changing requests (backend isSafeMethod complement).
	assertContains(t, content, "XSRF-TOKEN=([^;]*)")
	assertContains(t, content, "if (!['GET', 'HEAD', 'OPTIONS'].includes(request.method))")
	assertContains(t, content, "request.headers.set('X-XSRF-TOKEN', token)")

	// No bearer wiring: no store import, no Bearer header, no refresh flow.
	assertNotContains(t, content, "useAuthStore")
	assertNotContains(t, content, "Authorization")
	assertNotContains(t, content, "refreshInFlight")
	assertNotContains(t, content, "withAuthRetry")
	assertNotContains(t, content, "localStorage")

	// BUG-113 order guarantee (cookie mode): the 401 convergence middleware
	// (onResponse) runs first and passes the response through; the wrapper
	// then promotes r.error to a throw on the final result — both coexist.
	assertContains(t, content, "async onResponse({ response })")
	assertContains(t, content, ".then(r => { const d = r.data; const e = r.error; if (e !== undefined) throw e; return d as Res<'ListItems'> })")
	assertContains(t, content, ".then(r => { const d = r.data; const e = r.error; if (e !== undefined) throw e; return d as Res<'CreateItem'> })")
}
