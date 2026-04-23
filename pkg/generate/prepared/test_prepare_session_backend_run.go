//ff:func feature=generate type=test-helper control=sequence
//ff:what runPrepareSessionBackendCase — 단일 테이블 엔트리를 실행하는 클로저 생성

package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runPrepareSessionBackendCase returns a t.Run body for a single case.
// Keeping the assertion out of the outer range keeps each function
// within filefunc Q4 nesting/pure-line limits.
func runPrepareSessionBackendCase(tc prepareSessionBackendCase) func(*testing.T) {
	return func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: tc.manifest, ServiceFuncs: tc.funcs}
		got := sessionBackendFor(fs)
		assertPrepareSessionBackend(t, tc, got)
	}
}
