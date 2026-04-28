//ff:func feature=cli-init type=util control=iteration dimension=1
//ff:what createSkeletonDirs — creates every directory in the target tree

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"
)

// createSkeletonDirs creates every directory in the target tree. Ordering is
// from shallow to deep so os.MkdirAll never has to backfill intermediate
// nodes in later entries.
func createSkeletonDirs(targetDir string) error {
	for _, dir := range skeletonDirs() {
		abs := filepath.Join(targetDir, dir)
		if err := os.MkdirAll(abs, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", abs, err)
		}
	}
	return nil
}
