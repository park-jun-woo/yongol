//ff:func feature=gen-react type=test control=sequence
//ff:what assertNoRefreshStore — writeSessionStore 출력에 dead refresh 표면이 없는지 검증하는 테스트 헬퍼 (BUG-135)
package react

import (
	"os"
	"path/filepath"
	"testing"
)

// assertNoRefreshStore writes the bearer session store without a refresh
// signal and asserts the emitted auth.ts carries no refresh field, no setAuth
// refresh argument and no refresh reset in clear (BUG-135). Extracted from
// TestWriteSessionStore_NoRefreshDropsDeadCode so the table loop body stays
// within the per-block line limit.
func assertNoRefreshStore(t *testing.T, store string) {
	dir := t.TempDir()
	if err := writeSessionStore(dir, store, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "stores", "auth.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "token: string | null")
	assertContains(t, content, "setAuth: (token?: string | null) => void")
	assertContains(t, content, "setAuth: (token) => set({ token: token ?? null }),")
	assertContains(t, content, "clear: () => set({ token: null }),")
	// no dead refresh surface
	assertNotContains(t, content, "refresh")
}
