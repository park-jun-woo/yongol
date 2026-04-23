//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what removeStaleGoModFiles — 이전 run 의 go.mod/go.sum 잔재 제거

package gogin

import (
	"fmt"
	"os"
	"path/filepath"
)

// removeStaleGoModFiles sweeps go.mod / go.sum leftovers before a fresh
// `go mod init` run. `go mod init` refuses to overwrite an existing
// module file, and a leftover go.sum would keep entries from a previous
// dependency set alive.
func removeStaleGoModFiles(backendDir string) error {
	for _, name := range []string{"go.mod", "go.sum"} {
		p := filepath.Join(backendDir, name)
		if err := os.Remove(p); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove stale %s: %w", name, err)
		}
	}
	return nil
}
