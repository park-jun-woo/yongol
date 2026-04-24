//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what pruneOrphans — arts/tests/ 에만 남은 이전 .hurl 산출물 삭제 (drift 방지)

package hurl_mirror

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// pruneOrphans walks dstRoot and deletes any *.hurl file whose
// specs-relative path is not present in kept. The expectation is that
// kept tracks every file just mirrored, so the left-over set is exactly
// the stale artifacts from a previous generate run whose corresponding
// source has since been removed.
//
// Missing dstRoot is a no-op — the user has not run generate yet, so
// there is nothing to reconcile.
func pruneOrphans(dstRoot string, kept map[string]struct{}) error {
	info, err := os.Stat(dstRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", dstRoot, err)
	}
	if !info.IsDir() {
		return nil
	}
	return filepath.WalkDir(dstRoot, func(path string, d os.DirEntry, werr error) error {
		if werr != nil {
			return werr
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(d.Name()), ".hurl") {
			return nil
		}
		rel, rerr := filepath.Rel(dstRoot, path)
		if rerr != nil {
			return fmt.Errorf("rel %s: %w", path, rerr)
		}
		if _, ok := kept[rel]; ok {
			return nil
		}
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("remove orphan %s: %w", path, err)
		}
		return nil
	})
}
