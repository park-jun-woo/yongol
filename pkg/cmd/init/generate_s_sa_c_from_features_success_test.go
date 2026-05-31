//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateSSaCFromFeatures — 도메인별 stub 생성 / mkdir 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateSSaCFromFeatures_Success(t *testing.T) {
	target := t.TempDir()
	feats := []features.Feature{
		{Op: "CreateTask", Path: "POST /tasks"},
		{Op: "GetUser", Path: "GET /users/{id}"},
	}
	if err := generateSSaCFromFeatures(target, feats); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(target, "specs", "service", "task", "CreateTask.ssac"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "func CreateTask()") {
		t.Errorf("unexpected stub: %q", got)
	}
	if _, err := os.Stat(filepath.Join(target, "specs", "service", "user", "GetUser.ssac")); err != nil {
		t.Errorf("user stub missing: %v", err)
	}
}
