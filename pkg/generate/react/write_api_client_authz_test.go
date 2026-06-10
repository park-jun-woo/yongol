//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient bearer 미들웨어 포함 검증 — Bearer 토큰을 세션 store에서 읽는다

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_Bearer_AuthMiddleware(t *testing.T) {
	dir := t.TempDir()
	if err := writeAPIClient(dir, nil, true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "client.use({")
	assertContains(t, content, "async onRequest({ request })")
	assertContains(t, content, "import { useAuthStore } from '../stores/auth'")
	assertContains(t, content, "const token = useAuthStore.getState().token")
	assertContains(t, content, "request.headers.set('Authorization', `Bearer ${token}`)")
	assertNotContains(t, content, "localStorage")
}
