//ff:func feature=features type=test control=sequence
//ff:what TestRunAdd — new/existing 파싱 에러 / no-new-diff / 신규 op stub 생성 + skip 기존파일 분기 검증
package features

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAdd_CreatesStubAndHash(t *testing.T) {
	specs := t.TempDir()
	writeFeats(t, filepath.Join(specs, "features.yaml"), baseFeats)
	newContent := baseFeats + `  - op: DeleteTask
    path: DELETE /tasks/{id}
    desc: Delete a task
`
	newPath := filepath.Join(t.TempDir(), "new.yaml")
	writeFeats(t, newPath, newContent)

	var out bytes.Buffer
	if err := RunAdd(&out, specs, newPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// SSaC stub created under service/task/DeleteTask.ssac
	stub := filepath.Join(specs, "service", "task", "DeleteTask.ssac")
	if _, err := os.Stat(stub); err != nil {
		t.Fatalf("expected stub at %s: %v", stub, err)
	}
	// .yongol hash written.
	if _, err := os.Stat(filepath.Join(specs, ".yongol")); err != nil {
		t.Errorf("expected .yongol hash: %v", err)
	}
	// features.yaml replaced with new content.
	got, _ := os.ReadFile(filepath.Join(specs, "features.yaml"))
	if !strings.Contains(string(got), "DeleteTask") {
		t.Errorf("features.yaml not replaced: %q", got)
	}
	if !strings.Contains(out.String(), "1 new feature") {
		t.Errorf("want summary, got %q", out.String())
	}
}
