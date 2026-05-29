//ff:func feature=orchestrator type=test-helper control=sequence
//ff:what newTmpSpecsDir — tempdir 기반 specs 루트 생성 헬퍼
package yongol

import (
	"testing"
)

// newTmpSpecsDir creates an empty temp dir that tests can populate with
// partial SSOT layouts. Callers MkdirAll / WriteFile individually so that each
// test can assert a minimal matrix of presence states.
func newTmpSpecsDir(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}
