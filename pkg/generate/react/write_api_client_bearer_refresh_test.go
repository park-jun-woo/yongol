//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient bearer+refresh — single-flight refresh·재시도 1회·withAuthRetry 래핑 방출 검증 (auth-required op만 래핑)

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_Bearer_RefreshFlow(t *testing.T) {
	dir := t.TempDir()
	sec := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearer": {}}}
	optOut := openapi3.SecurityRequirements{}
	securedListItems := buildTokenOp("ListItems", []string{"items"}, nil)
	securedListItems.Security = &sec
	unsecuredRefresh := buildTokenOp("Refresh", []string{"access_token", "refresh_token"}, []string{"refresh_token"})
	unsecuredRefresh.Security = &optOut
	doc := &openapi3.T{
		Security: sec,
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: unsecuredRefresh}),
			openapi3.WithPath("/items", &openapi3.PathItem{Get: securedListItems}),
		),
	}
	plan := apiClientPlan{bearer: true, refresh: &refreshPlan{
		opID: "Refresh", method: "POST", path: "/auth/refresh",
		tokenField: "access_token", refreshField: "refresh_token", bodyKey: "refresh_token",
	}}
	if err := writeAPIClient(dir, doc, plan); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// Bearer injection stays; the immediate 401 logout middleware is gone.
	assertContains(t, content, "request.headers.set('Authorization', `Bearer ${token}`)")
	assertNotContains(t, content, "async onResponse")

	// Single-flight: module-level shared Promise, reset on settle.
	assertContains(t, content, "let refreshInFlight: Promise<boolean> | null = null")
	assertContains(t, content, "if (!refreshInFlight) {")
	assertContains(t, content, "refreshInFlight = null")

	// Refresh call: stored refresh token in the declared body key; store commit.
	assertContains(t, content, "client.POST('/auth/refresh', { body: { refresh_token: refresh } as any })")
	assertContains(t, content, "useAuthStore.getState().setAuth(data['access_token'], data['refresh_token'])")

	// Retry once, give up (clear + /login) on the second 401 or refresh failure.
	assertContains(t, content, "const retried = await call()")
	assertContains(t, content, "if (retried.response.status === 401) {")
	assertContains(t, content, "clearSessionAndLogin()")
	assertContains(t, content, "window.location.href = '/login'")

	// Auth-required ops are wrapped in withAuthRetry.
	assertContains(t, content, "withAuthRetry(() => client.GET('/items'")
	// Unsecured ops (Refresh) are NOT wrapped — a 401 on refresh is a
	// genuine failure, not a cue to retry (Phase046, BUG-146 Item 1).
	assertNotContains(t, content, "withAuthRetry(() => client.POST('/auth/refresh'")
	assertNotContains(t, content, "localStorage")
}
