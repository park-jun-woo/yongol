//ff:func feature=gen-hurl type=util control=sequence
//ff:what mirrorUserHurlFiles — specs/tests/scenario-*.hurl, invariant-*.hurl 을 arts/tests/ 로 verbatim 복사
package hurl

import (
	"fmt"
	"os"
	"path/filepath"
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
	return mirrorUserHurlEntries(srcDir, dstDir, entries)
}
