//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient cookie 모드 + csrf 비활성 — credentials include + 401 수렴만 방출, CSRF/store 없음 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_Cookie_CSRFDisabled(t *testing.T) {
	dir := t.TempDir()
	plan := apiClientPlan{cookie: true, csrf: false}
	if err := writeAPIClient(dir, nil, plan); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "credentials: 'include'")
	// Phase005: cookie mode always converges protected-fetch 401s to /login
	// (the optimistic ProtectedRoute delegates the auth decision here).
	assertContains(t, content, "if (response.status === 401 && window.location.pathname !== '/login')")
	assertContains(t, content, "window.location.href = '/login'")
	assertNotContains(t, content, "csrfToken")
	assertNotContains(t, content, "useAuthStore")
}
