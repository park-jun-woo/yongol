//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient bearer+refresh — single-flight refresh·재시도 1회·withAuthRetry 래핑 방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_Bearer_RefreshFlow(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: buildTokenOp("Refresh", []string{"access_token", "refresh_token"}, []string{"refresh_token"})}),
		openapi3.WithPath("/items", &openapi3.PathItem{Get: buildTokenOp("ListItems", []string{"items"}, nil)}),
	)}
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

	// Every operation entry is wrapped in withAuthRetry.
	assertContains(t, content, "withAuthRetry(() => client.GET('/items'")
	assertContains(t, content, "withAuthRetry(() => client.POST('/auth/refresh'")
	assertNotContains(t, content, "localStorage")
}
