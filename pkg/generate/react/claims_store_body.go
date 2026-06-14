//ff:func feature=gen-react type=generator control=sequence
//ff:what claimsStoreBody — claims store 상태 객체 본문 라인들 (setAuth 불변 + setClaim + clear 의 claims 소거)

package react

// claimsStoreBody returns the state-object body lines of the claims store
// (rendered at both persist and memory depths via indentLines). In the
// bearer shape setAuth is byte-identical to writeSessionStore's — it never
// touches claims (zustand's shallow merge preserves them across a 401
// refresh) — and clear() resets claims together with the tokens; hasRefresh
// (resolveHasRefresh) gates the refresh field/argument the same way
// writeSessionStore does (BUG-135). The cookie-mode claims-only shape keeps
// just claims/setClaim/clear (plans/stml/sitemap Phase005).
func claimsStoreBody(tokens, hasRefresh bool) []string {
	claimsOnly := []string{
		"claims: {},",
		"setClaim: (name, value) =>",
		"  set((state) => ({ claims: { ...state.claims, [name]: value } })),",
		"clear: () => set({ claims: {} }),",
	}
	if !tokens {
		return claimsOnly
	}
	if hasRefresh {
		return claimsStoreTokenRefreshBody()
	}
	return []string{
		"token: null,",
		"claims: {},",
		"setAuth: (token) => set({ token: token ?? null }),",
		"setClaim: (name, value) =>",
		"  set((state) => ({ claims: { ...state.claims, [name]: value } })),",
		"clear: () => set({ token: null, claims: {} }),",
	}
}
