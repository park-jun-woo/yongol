//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient authz 활성 시 401 자동 로그아웃 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_Authz_401AutoLogout(t *testing.T) {
	dir := t.TempDir()
	if err := writeAPIClient(dir, nil, true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "async onResponse({ response })")
	assertContains(t, content, "response.status === 401")
	assertContains(t, content, "localStorage.removeItem('access_token')")
	assertContains(t, content, "localStorage.removeItem('refresh_token')")
	assertContains(t, content, "window.location.href = '/login'")
}
