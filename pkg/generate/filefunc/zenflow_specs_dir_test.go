//ff:func feature=gen-filefunc type=test-helper control=iteration dimension=1
//ff:what zenflow dummy specs 디렉토리를 module 루트에서 탐색 (없으면 빈 문자열)
package filefunc

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// zenflowSpecsDir resolves dummys/zenflow/try-02/specs relative to the
// module root by walking up from the test file. Returns "" if not found.
func zenflowSpecsDir(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	dir := filepath.Dir(thisFile)
	for i := 0; i < 10; i++ {
		candidate := filepath.Join(dir, "dummys", "zenflow", "try-02", "specs")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
	return ""
}
