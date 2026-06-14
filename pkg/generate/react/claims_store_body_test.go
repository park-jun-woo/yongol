//ff:func feature=gen-react type=test control=sequence
//ff:what claimsStoreBody — bearer 전체형의 setAuth claims 불변·clear 소거 / cookie 축소형의 token 부재 검증

package react

import (
	"strings"
	"testing"
)

func TestClaimsStoreBody(t *testing.T) {
	full := strings.Join(claimsStoreBody(true, true), "\n")
	// bearer shape keeps the token fields and the writeSessionStore
	// setAuth contract — its body never touches claims, so the 401
	// refresh flow preserves them.
	assertContains(t, full, "token: null,")
	assertContains(t, full, "refresh: null,")
	assertContains(t, full, "setAuth: (token, refresh) =>")
	assertContains(t, full, "refresh: refresh === undefined ? state.refresh : refresh,")
	setAuthBody := full[strings.Index(full, "setAuth:"):strings.Index(full, "setClaim:")]
	assertNotContains(t, setAuthBody, "claims")
	// setClaim merges into the claims map; clear wipes claims with the tokens.
	assertContains(t, full, "set((state) => ({ claims: { ...state.claims, [name]: value } })),")
	assertContains(t, full, "clear: () => set({ token: null, refresh: null, claims: {} }),")

	reduced := strings.Join(claimsStoreBody(false, false), "\n")
	// cookie shape: claims/setClaim/clear only — no token surface at all.
	assertContains(t, reduced, "claims: {},")
	assertContains(t, reduced, "setClaim: (name, value) =>")
	assertContains(t, reduced, "clear: () => set({ claims: {} }),")
	assertNotContains(t, reduced, "token")
	assertNotContains(t, reduced, "setAuth")

	// bearer without refresh (BUG-135): token + claims kept, refresh dropped.
	noRefresh := strings.Join(claimsStoreBody(true, false), "\n")
	assertContains(t, noRefresh, "token: null,")
	assertContains(t, noRefresh, "claims: {},")
	assertContains(t, noRefresh, "setAuth: (token) => set({ token: token ?? null }),")
	assertContains(t, noRefresh, "clear: () => set({ token: null, claims: {} }),")
	assertNotContains(t, noRefresh, "refresh")
}
