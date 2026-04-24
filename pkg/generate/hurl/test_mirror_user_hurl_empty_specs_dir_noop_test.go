//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorUserHurlFiles_EmptySpecsDirNoop — 빈 specsDir 인자도 safe no-op

package hurl

import (
	"testing"
)

// TestMirrorUserHurlFiles_EmptySpecsDirNoop ensures empty specsDir
// argument is a safe no-op (defensive path for missing Fullstack).
func TestMirrorUserHurlFiles_EmptySpecsDirNoop(t *testing.T) {
	if err := mirrorUserHurlFiles("", t.TempDir()); err != nil {
		t.Fatalf("mirrorUserHurlFiles(\"\"): %v", err)
	}
}
