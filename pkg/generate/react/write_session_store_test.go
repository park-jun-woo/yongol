//ff:func feature=gen-react type=test control=sequence
//ff:what writeSessionStore — localStorage persist 래핑 / memory 비영속 방출 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSessionStore_LocalStoragePersist(t *testing.T) {
	dir := t.TempDir()
	if err := writeSessionStore(dir, "localStorage"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "stores", "auth.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { create } from 'zustand'")
	assertContains(t, content, "import { persist } from 'zustand/middleware'")
	assertContains(t, content, "export const useAuthStore = create<AuthState>()(")
	assertContains(t, content, "persist(")
	assertContains(t, content, "{ name: 'auth' }")
	assertContains(t, content, "token: string | null")
	assertContains(t, content, "refresh: string | null")
	assertContains(t, content, "setAuth: (token?: string | null, refresh?: string | null) => void")
	assertContains(t, content, "clear: () => set({ token: null, refresh: null })")
}
