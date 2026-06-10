//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient bearer 활성 시 401 → store 클리어 + /login 이동 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_Bearer_401AutoLogout(t *testing.T) {
	dir := t.TempDir()
	if err := writeAPIClient(dir, nil, apiClientPlan{bearer: true}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "async onResponse({ response })")
	assertContains(t, content, "response.status === 401")
	assertContains(t, content, "useAuthStore.getState().clear()")
	assertContains(t, content, "window.location.href = '/login'")
	assertNotContains(t, content, "localStorage")
}
