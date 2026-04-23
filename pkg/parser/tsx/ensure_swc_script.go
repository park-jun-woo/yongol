//ff:func feature=tsx-parser type=loader control=sequence
//ff:what ensureSwcScript — 번들된 swc parse.cjs 래퍼를 tmp 디렉토리에 1회만 기록

package tsx

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	swcScriptOnce sync.Once
	swcScriptPath string
	swcScriptErr  error
)

// ensureSwcScript writes the bundled parser wrapper into a deterministic
// temp file and returns its path. The script is tiny and stateless, so
// reuse across invocations is safe.
func ensureSwcScript() (string, error) {
	swcScriptOnce.Do(func() {
		dir := filepath.Join(os.TempDir(), "yongol-tsx-parser")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			swcScriptErr = fmt.Errorf("create swc script dir: %w", err)
			return
		}
		path := filepath.Join(dir, "parse.cjs")
		if err := os.WriteFile(path, []byte(swcParserScript), 0o644); err != nil {
			swcScriptErr = fmt.Errorf("write swc script: %w", err)
			return
		}
		swcScriptPath = path
	})
	return swcScriptPath, swcScriptErr
}
