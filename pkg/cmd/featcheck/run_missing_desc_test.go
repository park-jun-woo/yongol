//ff:func feature=cli-featcheck type=test control=sequence
//ff:what TestRunErrorBranches — read 실패 / parse 실패 / 빈 features / path·desc 누락 분기 검증
package featcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_MissingDesc(t *testing.T) {
	path := filepath.Join(t.TempDir(), "features.yaml")
	content := `features:
  - op: CreateTask
    path: POST /tasks
    desc: ""
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	_, _, err := Run(path)
	if err == nil || !strings.Contains(err.Error(), "missing required field 'desc'") {
		t.Fatalf("want missing desc error, got %v", err)
	}
}
