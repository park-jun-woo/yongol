//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient JWT 미들웨어 포함 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_Authz_JWTMiddleware(t *testing.T) {
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
	assertContains(t, content, "localStorage.getItem('access_token')")
	assertContains(t, content, "request.headers.set('Authorization', `Bearer ${token}`)")
}
