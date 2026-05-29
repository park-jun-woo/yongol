//ff:func feature=orchestrator type=loader control=sequence dimension=1
//ff:what parseYongolPkgSpecs missing-root 테스트 — pkgRoot 해석 실패 시 slog.Debug 만 남기고 Warn 은 발생시키지 않는다

package yongol

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestParseYongolPkgSpecsMissingRoot verifies that when findYongolPkgRoot
// cannot locate the built-in ssac/pkg tree (intentional in minimal CI
// containers), parseYongolPkgSpecs emits only a Debug log — not a Warn —
// since this is the expected "no built-ins available" state rather than a
// parse failure.
func TestParseYongolPkgSpecsMissingRoot(t *testing.T) {
	// Point the env override at a non-existent path so findYongolPkgRoot
	// falls through to the CWD / GOMODCACHE lookups, which we also suppress.
	t.Setenv("YONGOL_SSAC_PKG", filepath.Join(t.TempDir(), "does-not-exist"))
	// Isolate CWD + GOMODCACHE walkers from any ambient sibling ssac repo.
	isolated := t.TempDir()
	if err := os.Chdir(isolated); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Setenv("GOMODCACHE", filepath.Join(isolated, "empty-gomodcache"))

	if got := findYongolPkgRoot(); got != "" {
		t.Skipf("findYongolPkgRoot resolved %q despite isolation; skipping missing-root assertion", got)
	}

	buf := &bytes.Buffer{}
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(prev) })

	fs := &Fullstack{}
	parseYongolPkgSpecs(fs)

	if len(fs.YongolPkgSpecs) != 0 {
		t.Fatalf("expected no specs when pkg root missing; got %d", len(fs.YongolPkgSpecs))
	}
	log := buf.String()
	if strings.Contains(log, "level=WARN") {
		t.Fatalf("did not expect WARN log when pkg root missing; log=%q", log)
	}
	if !strings.Contains(log, "built-in pkg root not found") {
		t.Fatalf("expected debug log for missing pkg root; log=%q", log)
	}
}
