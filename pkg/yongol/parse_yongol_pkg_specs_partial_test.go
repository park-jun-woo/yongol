//ff:func feature=orchestrator type=loader control=sequence dimension=1
//ff:what parseYongolPkgSpecs partial success 테스트 — 일부 파일 실패 시 나머지는 로드하고 slog.Warn 을 emit 한다

package yongol

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestParseYongolPkgSpecsPartialSuccess verifies that a mix of valid and
// invalid Go files under the built-in pkg root results in (a) the valid
// funcspecs being loaded and (b) a slog.Warn event being emitted. This
// guarantees the silent-drop regression from parse_all.go stays fixed.
func TestParseYongolPkgSpecsPartialSuccess(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "auth")
	if err := os.MkdirAll(authDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	// Valid funcspec — should be loaded.
	ok := `package auth

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
	if err := os.WriteFile(filepath.Join(authDir, "hash_password.go"), []byte(ok), 0o644); err != nil {
		t.Fatalf("write ok: %v", err)
	}

	// Broken Go file — should cause a parse diagnostic but not block the
	// valid spec above.
	bad := `package auth

this is not valid Go source
`
	if err := os.WriteFile(filepath.Join(authDir, "broken.go"), []byte(bad), 0o644); err != nil {
		t.Fatalf("write bad: %v", err)
	}

	t.Setenv("YONGOL_SSAC_PKG", dir)

	buf := &bytes.Buffer{}
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(prev) })

	fs := &Fullstack{}
	parseYongolPkgSpecs(fs)

	if len(fs.YongolPkgSpecs) == 0 {
		t.Fatalf("expected partial success to keep valid specs; got 0")
	}
	log := buf.String()
	if !strings.Contains(log, "built-in funcspec load issues") {
		t.Fatalf("expected slog.Warn emit for load issues; log=%q", log)
	}
	if !strings.Contains(log, "pkg_root=") {
		t.Fatalf("expected pkg_root attr in log; log=%q", log)
	}
	if !strings.Contains(log, "diag_count=") {
		t.Fatalf("expected diag_count attr in log; log=%q", log)
	}
	// Ensure slog.Warn actually fired (level=WARN) and not some lesser level.
	if !strings.Contains(log, "level=WARN") {
		t.Fatalf("expected WARN level; log=%q", log)
	}
	_ = context.TODO()
}
