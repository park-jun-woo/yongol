//ff:func feature=cli-init type=util control=sequence
//ff:what ensureEmptyDir — refuse to touch a directory that already contains files unless --force is set

package cliinit

import (
	"fmt"
	"os"
)

// ensureEmptyDir refuses to touch a directory that already contains files
// unless --force is set. A dedicated helper keeps Run() readable and makes
// the semantic easy to unit-test.
func ensureEmptyDir(dir string, force bool) error {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", dir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s exists but is not a directory", dir)
	}
	if force {
		return nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read %s: %w", dir, err)
	}
	if len(entries) > 0 {
		return fmt.Errorf("%s is not empty (use -f/--force to override)", dir)
	}
	return nil
}
