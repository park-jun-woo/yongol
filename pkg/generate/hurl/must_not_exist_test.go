//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what mustNotExist — 파일이 존재하지 않는지 검증

package hurl

import (
	"os"
	"testing"
)

// mustNotExist fails the test if the file at path exists.
func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected %s to not exist; err=%v", path, err)
	}
}
