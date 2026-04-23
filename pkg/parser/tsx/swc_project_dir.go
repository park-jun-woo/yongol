//ff:func feature=tsx-parser type=loader control=iteration dimension=1
//ff:what swcProjectDir — node 실행 디렉토리 결정 (env override → node_modules 조상 → package.json 조상)

package tsx

import (
	"os"
	"path/filepath"
)

// swcProjectDir returns the directory node should run from so that a local
// @swc/core install is discovered. The resolution order is:
//
//  1. $YONGOL_SWC_PROJECT_DIR (explicit override)
//  2. ascending from the target file: nearest ancestor with node_modules/@swc/core
//  3. ascending from the target file: nearest ancestor with package.json
//  4. the file's own directory
//
// The caller is responsible for a clear diagnostic when none of the above
// resolves to a working install; swcRunnerNotice is the canonical message.
func swcProjectDir(target string) string {
	if v := os.Getenv("YONGOL_SWC_PROJECT_DIR"); v != "" {
		return v
	}
	abs, err := filepath.Abs(target)
	if err != nil {
		abs = target
	}
	dir := filepath.Dir(abs)
	// Walk up to filesystem root.
	for {
		if st, err := os.Stat(filepath.Join(dir, "node_modules", "@swc", "core")); err == nil && st.IsDir() {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	// Second pass: nearest package.json.
	dir = filepath.Dir(abs)
	for {
		if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return filepath.Dir(abs)
}
