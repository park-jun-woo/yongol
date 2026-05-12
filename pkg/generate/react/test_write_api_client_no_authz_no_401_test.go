//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient authz 미설정 시 401 핸들러 미방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_NoAuthz_No401Handler(t *testing.T) {
	dir := t.TempDir()
	if err := writeAPIClient(dir, nil, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertNotContains(t, content, "onResponse")
	assertNotContains(t, content, "response.status === 401")
	assertNotContains(t, content, "window.location.href")
}
