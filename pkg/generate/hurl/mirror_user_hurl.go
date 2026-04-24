//ff:func feature=gen-hurl type=util control=sequence
//ff:what mirrorUserHurlFiles — specs/tests/scenario-*.hurl, invariant-*.hurl 을 arts/tests/ 로 verbatim 복사
package hurl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// mirrorUserHurlFiles copies user-authored scenario-*.hurl and
// invariant-*.hurl from specsDir/tests/ into artifactsDir/tests/.
//
// Scope:
//   - Only files whose base name matches "scenario-*.hurl" or
//     "invariant-*.hurl" are mirrored. The auto-generated smoke.hurl
//     (written by writeSmokeHurl) is skipped even if present in specs
//     to avoid overwriting the generator's output.
//   - Missing specsDir or specsDir/tests/ is a silent no-op — projects
//     without user-authored hurl still get the auto-gen smoke.hurl.
//   - Files are copied verbatim (no transformation).
//
// Rationale — docs/scenario.md assumes `arts/<project>/tests/*.hurl`
// is the canonical run target, so the generator must own that mirror.
// Without it smoke.hurl is the only runnable file and user scenarios
// are silently dropped (BUG-026).
func mirrorUserHurlFiles(specsDir, artifactsDir string) error {
	if specsDir == "" {
		return nil
	}
	srcDir := filepath.Join(specsDir, "tests")
	info, err := os.Stat(srcDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", srcDir, err)
	}
	if !info.IsDir() {
		return nil
	}
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("readdir %s: %w", srcDir, err)
	}
	dstDir := filepath.Join(artifactsDir, "tests")
	var mirrored bool
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !isUserHurlName(name) {
			continue
		}
		if !mirrored {
			if err := os.MkdirAll(dstDir, 0o755); err != nil {
				return fmt.Errorf("mkdir %s: %w", dstDir, err)
			}
			mirrored = true
		}
		src := filepath.Join(srcDir, name)
		dst := filepath.Join(dstDir, name)
		if err := copyHurlFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

// isUserHurlName reports whether name is a user-authored hurl file
// that should be mirrored. smoke.hurl is explicitly excluded so the
// auto-generator never clashes with a stray user file of the same
// name.
func isUserHurlName(name string) bool {
	if !strings.HasSuffix(name, ".hurl") {
		return false
	}
	if name == "smoke.hurl" {
		return false
	}
	return strings.HasPrefix(name, "scenario-") || strings.HasPrefix(name, "invariant-")
}

// copyHurlFile copies src → dst verbatim. Truncates dst if it exists.
func copyHurlFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %s: %w", src, err)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create %s: %w", dst, err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s -> %s: %w", src, dst, err)
	}
	return nil
}
