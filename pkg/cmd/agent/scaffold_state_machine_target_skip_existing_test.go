//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScaffoldStateMachineTargetSkipExisting(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "orders.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	created, err := scaffoldStateMachineTarget(dir, "orders", []string{"pending"}, nil, "sys", Config{}, &out)
	if err != nil {
		t.Fatalf("skip-existing: unexpected error: %v", err)
	}
	if created {
		t.Fatal("skip-existing: expected created=false")
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}
