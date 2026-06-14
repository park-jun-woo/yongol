//ff:func feature=gen-react type=test control=sequence
//ff:what writeSessionStoreClaims — bearer 전체형/cookie claims 전용형 × persist/memory 방출과 setAuth 불변·clear 의 claims 소거 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteSessionStoreClaims(t *testing.T) {
	read := func(t *testing.T, dir string) string {
		t.Helper()
		data, err := os.ReadFile(filepath.Join(dir, "stores", "auth.ts"))
		if err != nil {
			t.Fatalf("read store: %v", err)
		}
		return string(data)
	}

	t.Run("bearer localStorage keeps tokens and adds claims", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeSessionStoreClaims(dir, "localStorage", true, true); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertContains(t, out, "import { persist } from 'zustand/middleware'")
		assertContains(t, out, "token: string | null")
		assertContains(t, out, "claims: Record<string, string>")
		assertContains(t, out, "setClaim: (name: string, value: string) => void")
		assertContains(t, out, "setClaim: (name, value) =>")
		assertContains(t, out, "set((state) => ({ claims: { ...state.claims, [name]: value } })),")
		// setAuth keeps the exact pre-claims contract — the refresh flow
		// calls setAuth, so claims survive a 401 refresh untouched.
		assertContains(t, out, "setAuth: (token, refresh) =>")
		assertContains(t, out, "refresh: refresh === undefined ? state.refresh : refresh,")
		// setAuth's body must not mention claims at all (preservation).
		setAuthBody := out[strings.Index(out, "setAuth: (token, refresh)"):strings.Index(out, "setClaim: (name, value)")]
		assertNotContains(t, setAuthBody, "claims")
		// clear() resets claims together with the tokens (logout wipes the
		// role so the menu hides immediately).
		assertContains(t, out, "clear: () => set({ token: null, refresh: null, claims: {} }),")
	})

	t.Run("bearer memory skips persist", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeSessionStoreClaims(dir, "memory", true, true); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertNotContains(t, out, "persist")
		assertContains(t, out, "claims: {},")
		assertContains(t, out, "clear: () => set({ token: null, refresh: null, claims: {} }),")
	})

	t.Run("bearer without refresh signal drops the dead refresh surface", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeSessionStoreClaims(dir, "localStorage", true, false); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		// token + claims survive; refresh is gone (BUG-135).
		assertContains(t, out, "token: string | null")
		assertContains(t, out, "claims: Record<string, string>")
		assertContains(t, out, "setAuth: (token?: string | null) => void")
		assertContains(t, out, "setAuth: (token) => set({ token: token ?? null }),")
		assertContains(t, out, "clear: () => set({ token: null, claims: {} }),")
		assertNotContains(t, out, "refresh")
	})

	t.Run("cookie localStorage emits the claims-only reduced store", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeSessionStoreClaims(dir, "localStorage", false, false); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertContains(t, out, "import { persist } from 'zustand/middleware'")
		assertContains(t, out, "claims: Record<string, string>")
		assertContains(t, out, "clear: () => set({ claims: {} }),")
		// No token surface at all — httpOnly cookies own the session.
		assertNotContains(t, out, "token")
		assertNotContains(t, out, "setAuth")
	})

	t.Run("cookie memory emits the claims-only store without persist", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeSessionStoreClaims(dir, "memory", false, false); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertNotContains(t, out, "persist")
		assertNotContains(t, out, "setAuth")
		assertContains(t, out, "clear: () => set({ claims: {} }),")
	})
}
