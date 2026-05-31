//ff:func feature=chain type=test control=iteration dimension=1
//ff:what TestChain — Chain 의 nil-OpenAPI / not-found / matched-ServiceFunc 분기 검증
package chain

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func opus47SpecsDir(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	dir := filepath.Dir(thisFile)
	for i := 0; i < 10; i++ {
		candidate := filepath.Join(dir, "examples", "zenflow", "opus4_7", "specs")
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
