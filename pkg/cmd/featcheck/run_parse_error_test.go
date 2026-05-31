//ff:func feature=cli-featcheck type=test control=sequence
//ff:what TestRunErrorBranches — read 실패 / parse 실패 / 빈 features / path·desc 누락 분기 검증
package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_ParseError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(path, []byte("features: [unterminated"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, err := Run(path)
	if err == nil || !strings.Contains(err.Error(), "parse features") {
		t.Fatalf("want parse error, got %v", err)
	}
}
