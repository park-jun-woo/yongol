//ff:func feature=gen-hurl type=command control=sequence
//ff:what MirrorSpecsTests — specs/tests/ → arts/tests/ 전체 1:1 디렉토리 미러링

package hurl_mirror

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MirrorSpecsTests copies every *.hurl file under specsDir/tests/ into
// artsDir/tests/, preserving sub-directory layout. It is the single entry
// point replacing the deleted pkg/generate/hurl/ emit engine; hurl files
// are entirely user-owned, so yongol only mirrors them byte-for-byte.
//
// Behaviour:
//   - specsDir/tests/ missing → no-op (no error). The H-2 WARNING at
//     validate time already surfaces the empty-tests case.
//   - specsDir/tests/ present → every *.hurl file is copied verbatim.
//   - Orphaned *.hurl files in artsDir/tests/ that no longer exist under
//     specsDir/tests/ are deleted to prevent drift (see pruneOrphans).
//   - Write failure → ERROR (propagated as-is).
//
// Returns the number of files mirrored so callers can surface a summary
// ("mirrored N hurl files").
func MirrorSpecsTests(specsDir, artsDir string) (int, error) {
	if specsDir == "" {
		return 0, nil
	}
	srcRoot := filepath.Join(specsDir, "tests")
	info, err := os.Stat(srcRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("stat %s: %w", srcRoot, err)
	}
	if !info.IsDir() {
		return 0, nil
	}
	dstRoot := filepath.Join(artsDir, "tests")

	mirrored := 0
	kept := map[string]struct{}{}
	walkErr := filepath.WalkDir(srcRoot, func(path string, d os.DirEntry, werr error) error {
		if werr != nil {
			return werr
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(d.Name()), ".hurl") {
			return nil
		}
		rel, rerr := filepath.Rel(srcRoot, path)
		if rerr != nil {
			return fmt.Errorf("rel %s: %w", path, rerr)
		}
		dst := filepath.Join(dstRoot, rel)
		if err := copyHurlFile(path, dst); err != nil {
			return err
		}
		kept[rel] = struct{}{}
		mirrored++
		return nil
	})
	if walkErr != nil {
		return mirrored, walkErr
	}
	if err := pruneOrphans(dstRoot, kept); err != nil {
		return mirrored, err
	}
	return mirrored, nil
}
