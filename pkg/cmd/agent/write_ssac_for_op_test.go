//ff:func feature=agent type=test control=sequence
//ff:what TestWriteSSaCForOp — operationId의 SSaC 파일 내용 기록, 부재 시 무기록 검증

package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteSSaCForOp(t *testing.T) {
	dir := t.TempDir()
	svc := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(svc, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(svc, "Login.ssac"), []byte("func Login() {}"), 0o644); err != nil {
		t.Fatal(err)
	}

	var b strings.Builder
	writeSSaCForOp(&b, dir, "Login")
	out := b.String()
	if !strings.Contains(out, "SSaC (Login.ssac):") || !strings.Contains(out, "func Login()") {
		t.Errorf("ssac → %q", out)
	}

	// Missing op: nothing.
	var b2 strings.Builder
	writeSSaCForOp(&b2, dir, "Missing")
	if b2.Len() != 0 {
		t.Errorf("missing op wrote %q", b2.String())
	}
}
