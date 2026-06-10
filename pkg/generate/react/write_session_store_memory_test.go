//ff:func feature=gen-react type=test control=sequence
//ff:what writeSessionStore memory 모드 — persist 없는 비영속 store 방출 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSessionStore_MemoryNonPersistent(t *testing.T) {
	dir := t.TempDir()
	if err := writeSessionStore(dir, "memory"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "stores", "auth.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { create } from 'zustand'")
	assertNotContains(t, content, "persist")
	assertContains(t, content, "export const useAuthStore = create<AuthState>()((set) => ({")
}
