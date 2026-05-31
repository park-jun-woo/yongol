//ff:func feature=cli type=test control=sequence
//ff:what TestManualForAIPath — manual-for-ai.md 발견(경로) / 미발견(github URL) 분기 검증
package main

import (
	"os"
	"testing"
)

func restoreCwd(t *testing.T) {
	t.Helper()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(orig) })
}
