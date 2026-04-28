//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what mustExist — 지정 경로에 기대한 내용이 기록되어 있는지 검증

package hurl_mirror

import (
	"os"
	"testing"
)

// mustExist reads path and fatals when the content differs from want.
func mustExist(t *testing.T, path, want string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if string(got) != want {
		t.Fatalf("content of %s = %q; want %q", path, string(got), want)
	}
}
