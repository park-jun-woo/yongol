//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient 산출 api.ts 의 baseUrl 이 빈 값인지 검증 (BUG-110)

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAPIClient_BaseUrlEmpty(t *testing.T) {
	dir := t.TempDir()
	if err := writeAPIClient(dir, nil, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "baseUrl: ''")
	assertNotContains(t, content, "baseUrl: '/api'")
}
