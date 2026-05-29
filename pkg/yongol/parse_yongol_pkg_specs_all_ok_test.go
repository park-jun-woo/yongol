//ff:func feature=orchestrator type=loader control=sequence dimension=1
//ff:what parseYongolPkgSpecs all-ok 테스트 — 모든 파일이 정상이면 slog.Warn 을 발생시키지 않는다

package yongol

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestParseYongolPkgSpecsAllOK verifies that when every file under the
// built-in pkg root parses cleanly, parseYongolPkgSpecs loads all specs and
// emits no WARN-level log. This guards against noisy operator logs in the
// common case.
func TestParseYongolPkgSpecsAllOK(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "auth")
	if err := os.MkdirAll(authDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	src := `package auth

// @func hashPassword
// @description hash

type HashPasswordRequest struct {
	Password string
}
type HashPasswordResponse struct {
	HashedPassword string
}
func HashPassword(req HashPasswordRequest) (HashPasswordResponse, error) {
	return HashPasswordResponse{HashedPassword: "hashed"}, nil
}
`
	if err := os.WriteFile(filepath.Join(authDir, "hash_password.go"), []byte(src), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	t.Setenv("YONGOL_SSAC_PKG", dir)

	buf := &bytes.Buffer{}
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(prev) })

	fs := &Fullstack{}
	parseYongolPkgSpecs(fs)

	if len(fs.YongolPkgSpecs) == 0 {
		t.Fatalf("expected specs to load on all-ok path")
	}
	log := buf.String()
	if strings.Contains(log, "built-in funcspec load issues") {
		t.Fatalf("did not expect load-issue warning on clean path; log=%q", log)
	}
	if strings.Contains(log, "level=WARN") {
		t.Fatalf("did not expect any WARN log on clean path; log=%q", log)
	}
}
