//ff:func feature=contract type=walker control=sequence
//ff:what CollectPreserved — 디렉토리를 walk 해 preserve 상태인 .go 파일 경로 목록을 반환

package contract

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// CollectPreserved walks rootDir recursively and returns the paths of
// every `.go` file whose DetectPreserved result is StatePreserved.
//
// Non-Go files, directories starting with "." (e.g. `.git`), and
// vendored paths (`vendor/`) are skipped up-front so callers do not
// have to pre-filter. Errors encountered while reading individual
// files are swallowed — a file that cannot be opened is treated as
// "not applicable" rather than aborting the scan.
func CollectPreserved(rootDir string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if name == "." {
				return nil
			}
			if strings.HasPrefix(name, ".") && path != rootDir {
				return filepath.SkipDir
			}
			if name == "vendor" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		state, derr := DetectPreserved(path)
		if derr != nil {
			return nil
		}
		if state == StatePreserved {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}
