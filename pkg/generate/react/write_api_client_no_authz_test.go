//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient JWT 미들웨어 미포함 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_NoAuthz_NoMiddleware(t *testing.T) {
	dir := t.TempDir()
	if err := writeAPIClient(dir, nil, apiClientPlan{}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertNotContains(t, content, "client.use")
	assertNotContains(t, content, "onRequest")
	assertNotContains(t, content, "Authorization")
}
