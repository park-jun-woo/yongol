//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what mustNotExist — 지정 경로가 존재하지 않는지 검증

package hurl_mirror

import (
	"os"
	"testing"
)

// mustNotExist fatals when path exists. Used to assert that pruning or
// no-op short-circuits did not create an artifacts directory.
func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected %s to not exist; stat err=%v", path, err)
	}
}
