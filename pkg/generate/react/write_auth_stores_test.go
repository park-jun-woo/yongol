//ff:func feature=gen-react type=test control=sequence
//ff:what writeAuthStores — claims 캡처 시 claims store(빈 store 는 localStorage 폴백)/bearer 는 현행 store/그 외 미방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAuthStores(t *testing.T) {
	claimsPages := []stml.PageSpec{{Actions: []stml.ActionBlock{{
		Captures: []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}},
	}}}}
	read := func(t *testing.T, dir string) string {
		t.Helper()
		data, err := os.ReadFile(filepath.Join(dir, "stores", "auth.ts"))
		if err != nil {
			t.Fatalf("read store: %v", err)
		}
		return string(data)
	}

	t.Run("claims captures emit the claims store, empty store falls back to localStorage", func(t *testing.T) {
		dir := t.TempDir()
		// authStore "" — claims captures without backend.auth resolved none
		if err := writeAuthStores(dir, "", false, false, claimsPages); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertContains(t, out, "setClaim: (name, value) =>")
		// the localStorage fallback wraps the store in zustand persist
		assertContains(t, out, "import { persist } from 'zustand/middleware'")
		// cookie mode (bearerAuth false) → claims-only reduced shape
		assertNotContains(t, out, "setAuth")
	})

	t.Run("claims captures in bearer mode keep the full token shape", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeAuthStores(dir, "memory", true, true, claimsPages); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertContains(t, out, "setAuth: (token, refresh) =>")
		assertContains(t, out, "setClaim: (name, value) =>")
		assertNotContains(t, out, "persist")
	})

	t.Run("bearer without claims captures keeps the pre-Phase005 store", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeAuthStores(dir, "localStorage", true, true, nil); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertContains(t, out, "setAuth: (token, refresh) =>")
		assertNotContains(t, out, "setClaim")
	})

	t.Run("bearer without refresh signal drops the dead refresh surface", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeAuthStores(dir, "localStorage", true, false, nil); err != nil {
			t.Fatal(err)
		}
		out := read(t, dir)
		assertContains(t, out, "setAuth: (token) => set({ token: token ?? null }),")
		assertNotContains(t, out, "setClaim")
		assertNotContains(t, out, "refresh")
	})

	t.Run("cookie or no-auth without claims emits no store", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeAuthStores(dir, "localStorage", false, false, nil); err != nil {
			t.Fatal(err)
		}
		if _, err := os.Stat(filepath.Join(dir, "stores", "auth.ts")); !os.IsNotExist(err) {
			t.Fatalf("expected no store file, stat err = %v", err)
		}
	})
}
