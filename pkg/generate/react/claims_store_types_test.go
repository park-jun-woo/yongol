//ff:func feature=gen-react type=test control=sequence
//ff:what claimsStoreTypes — bearer 전체형/cookie 축소형 AuthState interface 선언 검증

package react

import "testing"

func TestClaimsStoreTypes(t *testing.T) {
	full := claimsStoreTypes(true)
	// bearer shape: token/refresh/setAuth of writeSessionStore plus claims/setClaim
	assertContains(t, full, "interface AuthState {")
	assertContains(t, full, "token: string | null")
	assertContains(t, full, "refresh: string | null")
	assertContains(t, full, "claims: Record<string, string>")
	assertContains(t, full, "setAuth: (token?: string | null, refresh?: string | null) => void")
	assertContains(t, full, "setClaim: (name: string, value: string) => void")
	assertContains(t, full, "clear: () => void")

	reduced := claimsStoreTypes(false)
	// cookie shape: claims-only — no token, no refresh, no setAuth
	assertContains(t, reduced, "interface AuthState {")
	assertContains(t, reduced, "claims: Record<string, string>")
	assertContains(t, reduced, "setClaim: (name: string, value: string) => void")
	assertContains(t, reduced, "clear: () => void")
	assertNotContains(t, reduced, "token")
	assertNotContains(t, reduced, "setAuth")
}
