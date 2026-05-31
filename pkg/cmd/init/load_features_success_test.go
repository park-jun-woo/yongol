//ff:func feature=cli-init type=test control=sequence
//ff:what TestLoadFeatures — featcheck 에러 / ERROR 진단 / 정상 로드 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFeatures_Success(t *testing.T) {
	p := filepath.Join(t.TempDir(), "features.yaml")
	content := `tables:
  tasks: {}
features:
  - op: CreateTask
    path: POST /tasks
    desc: Create
    table: tasks
`
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	ff, err := loadFeatures(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ff == nil || len(ff.Features) != 1 {
		t.Errorf("unexpected ff: %+v", ff)
	}
}
