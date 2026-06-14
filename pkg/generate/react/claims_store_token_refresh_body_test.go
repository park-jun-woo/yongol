//ff:func feature=gen-react type=test control=sequence
//ff:what claimsStoreTokenRefreshBody — bearer+refresh claims store 본문이 token/refresh/claims·setAuth·clear 라인을 방출하는지 검증 (BUG-135)
package react

import (
	"strings"
	"testing"
)

func TestClaimsStoreTokenRefreshBody(t *testing.T) {
	body := strings.Join(claimsStoreTokenRefreshBody(), "\n")
	assertContains(t, body, "token: null,")
	assertContains(t, body, "refresh: null,")
	assertContains(t, body, "claims: {},")
	assertContains(t, body, "setAuth: (token, refresh) =>")
	assertContains(t, body, "refresh: refresh === undefined ? state.refresh : refresh,")
	assertContains(t, body, "set((state) => ({ claims: { ...state.claims, [name]: value } })),")
	assertContains(t, body, "clear: () => set({ token: null, refresh: null, claims: {} }),")
}
