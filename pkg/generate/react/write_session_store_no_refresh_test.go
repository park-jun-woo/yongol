//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what writeSessionStore — refresh 신호 없을 때 persist/memory 양쪽에서 dead refresh 표면(필드/인자/clear 리셋) 미방출 검증 (BUG-135)
package react

import "testing"

// TestWriteSessionStore_NoRefreshDropsDeadCode covers BUG-135: when no refresh
// signal exists (resolveHasRefresh false), the store emits no refresh field,
// no setAuth refresh argument and no refresh reset in clear — across both the
// localStorage-persist and memory emission depths.
func TestWriteSessionStore_NoRefreshDropsDeadCode(t *testing.T) {
	for _, store := range []string{"localStorage", "memory"} {
		t.Run(store, func(t *testing.T) { assertNoRefreshStore(t, store) })
	}
}
