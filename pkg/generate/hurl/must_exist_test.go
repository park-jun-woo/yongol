//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what mustExist — 파일이 존재하고 본문이 일치하는지 검증

package hurl

import (
	"os"
	"testing"
)

// mustExist fails the test if the file at path does not exist or the
// body does not match wantBody.
func mustExist(t *testing.T, path, wantBody string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if string(got) != wantBody {
		t.Fatalf("%s body = %q; want %q", path, string(got), wantBody)
	}
}
