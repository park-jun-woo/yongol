//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what writeFile — 테스트용 파일 쓰기 헬퍼

package hurl

import (
	"os"
	"testing"
)

// writeFile writes body to path and fails the test on error.
func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
