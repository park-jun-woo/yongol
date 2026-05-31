//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdMissingDir — MissingDir 서브테스트
package main

import (
	"testing"
)

func subtestTestGenerateCmdMissingDir(t *testing.T) {

	_, _, err := runCmd(t, "generate", "/tmp/nonexistent-yongol-specs", "/tmp/output")
	if err == nil {
		t.Fatal("expected error for missing dir, got nil")
	}

}
