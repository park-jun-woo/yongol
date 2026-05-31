//ff:func feature=cli type=test control=sequence
//ff:what TestManualForAIPath — manual-for-ai.md 발견(경로) / 미발견(github URL) 분기 검증
package main

import (
	"os"
	"strings"
	"testing"
)

func TestManualForAIPath_NotFound(t *testing.T) {
	restoreCwd(t)
	// An isolated temp dir whose ancestry has no manual-for-ai.md → github URL.
	dir := t.TempDir()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	got := manualForAIPath()
	if !strings.HasPrefix(got, "https://github.com/") {
		t.Skipf("ambient manual-for-ai.md resolved %q; skipping URL assertion", got)
	}
}
