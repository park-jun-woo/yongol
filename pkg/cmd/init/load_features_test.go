//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestLoadFeatures — featcheck 에러 / ERROR 진단 / 정상 로드 분기 검증

package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFeatures_FeatcheckError(t *testing.T) {
	// Missing file -> featcheck.Run returns an error, propagated.
	if _, err := loadFeatures(filepath.Join(t.TempDir(), "nope.yaml")); err == nil {
		t.Fatal("want error for missing features.yaml")
	}
}

func TestLoadFeatures_ValidationError(t *testing.T) {
	// Duplicate op -> FT-01 ERROR diagnostic -> validation failure.
	p := filepath.Join(t.TempDir(), "features.yaml")
	content := `tables:
  tasks: {}
features:
  - op: CreateTask
    path: POST /tasks
    desc: Create
    table: tasks
  - op: CreateTask
    path: POST /tasks/v2
    desc: Create v2
    table: tasks
`
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := loadFeatures(p)
	if err == nil || !strings.Contains(err.Error(), "validation failed") {
		t.Fatalf("want validation failure, got %v", err)
	}
}

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
