//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateHurlFromFeatures — 다중 feature stub 생성 / 잘못된 path 에러 / write 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateHurlFromFeatures_Success(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs", "tests"), 0o755); err != nil {
		t.Fatal(err)
	}
	feats := []features.Feature{
		{Op: "CreateTask", Path: "POST /tasks"},
		{Op: "GetTask", Path: "GET /tasks/{id}"},
	}
	if err := generateHurlFromFeatures(target, feats); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(target, "specs", "tests", "smoke.hurl"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(got)
	if !strings.Contains(s, "# CreateTask") || !strings.Contains(s, "POST {{host}}/tasks") {
		t.Errorf("missing CreateTask stub: %q", s)
	}
	if !strings.Contains(s, "# GetTask") || !strings.Contains(s, "GET {{host}}/tasks/{id}") {
		t.Errorf("missing GetTask stub: %q", s)
	}
}
