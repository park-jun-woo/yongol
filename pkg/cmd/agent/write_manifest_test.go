//ff:func feature=agent type=test control=sequence
//ff:what TestFixFileGroup — stall 추적(증가/리셋/3회skip) + fixFile 성공/실패 + ruleID hint 분기 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func writeManifest(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata: {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}
