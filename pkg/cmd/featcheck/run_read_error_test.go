//ff:func feature=cli-featcheck type=test control=sequence
//ff:what TestRunErrorBranches — read 실패 / parse 실패 / 빈 features / path·desc 누락 분기 검증
package featcheck

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_ReadError(t *testing.T) {
	_, _, err := Run(filepath.Join(t.TempDir(), "does-not-exist.yaml"))
	if err == nil || !strings.Contains(err.Error(), "read features") {
		t.Fatalf("want read error, got %v", err)
	}
}
