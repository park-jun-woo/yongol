//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient cookie 모드 + csrf 비활성 — credentials include 만 방출, 미들웨어 없음 검증

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
	assertNotContains(t, content, "client.use")
	assertNotContains(t, content, "csrfToken")
	assertNotContains(t, content, "useAuthStore")
}
