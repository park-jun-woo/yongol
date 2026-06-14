//ff:func feature=gen-react type=test control=sequence
//ff:what refresh 흐름 × claims store 회귀 — refresh 성공 시 setAuth 만 호출(claims 보존), 실패 시 clear() 가 claims 소거 고정

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRefreshFlowClaimsRegression pins the BUG-118-adjacent contract of
// plans/stml/sitemap Phase005: the 401 refresh flow and the claims store
// share src/stores/auth.ts, so the snapshots fix (1) a successful refresh
// commits through setAuth only — whose implementation never touches
// claims, preserving the role across token rotation — and (2) the refresh
// failure path calls clear(), whose implementation resets claims, so a
// dead session never keeps a stale role-gated menu visible.
func TestRefreshFlowClaimsRegression(t *testing.T) {
	// (1) the emitted refresh flow commits via setAuth and ends sessions
	// via clear() — exactly the two store functions whose claim behavior
	// the store snapshot below pins.
	var b strings.Builder
	writeRefreshFlow(&b, &refreshPlan{
		opID: "Refresh", method: "post", path: "/auth/refresh",
		tokenField: "access_token", refreshField: "refresh_token", bodyKey: "refresh_token",
	})
	flow := b.String()
	assertContains(t, flow, "useAuthStore.getState().setAuth(data['access_token'], data['refresh_token'])")
	assertContains(t, flow, "useAuthStore.getState().clear()")
	assertNotContains(t, flow, "setClaim") // refresh never rewrites claims

	// (2) the bearer claims store: setAuth's implementation is claims-free
	// (preservation) and clear()'s resets claims (eviction).
	dir := t.TempDir()
	if err := writeSessionStoreClaims(dir, "localStorage", true, true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "stores", "auth.ts"))
	if err != nil {
		t.Fatal(err)
	}
	store := string(data)
	setAuthImpl := store[strings.Index(store, "setAuth: (token, refresh)"):strings.Index(store, "setClaim: (name, value)")]
	assertNotContains(t, setAuthImpl, "claims")
	assertContains(t, store, "clear: () => set({ token: null, refresh: null, claims: {} }),")
}
