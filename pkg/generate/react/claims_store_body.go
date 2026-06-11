//ff:func feature=gen-react type=generator control=sequence
//ff:what claimsStoreBody — claims store 상태 객체 본문 라인들 (setAuth 불변 + setClaim + clear 의 claims 소거)

package react

// claimsStoreBody returns the state-object body lines of the claims store
// (rendered at both persist and memory depths via indentLines). In the
// bearer shape setAuth is byte-identical to writeSessionStore's — it never
// touches claims, so the 401 refresh flow preserves them — and clear()
// resets claims together with the tokens; the cookie-mode claims-only
// shape keeps just claims/setClaim/clear (plans/stml/sitemap Phase005).
func claimsStoreBody(tokens bool) []string {
	full := []string{
		"token: null,",
		"refresh: null,",
		"claims: {},",
		"setAuth: (token, refresh) =>",
		"  set((state) => ({",
		"    token: token ?? null,",
		"    refresh: refresh === undefined ? state.refresh : refresh,",
		"  })),",
		"setClaim: (name, value) =>",
		"  set((state) => ({ claims: { ...state.claims, [name]: value } })),",
		"clear: () => set({ token: null, refresh: null, claims: {} }),",
	}
	claimsOnly := []string{
		"claims: {},",
		"setClaim: (name, value) =>",
		"  set((state) => ({ claims: { ...state.claims, [name]: value } })),",
		"clear: () => set({ claims: {} }),",
	}
	if tokens {
		return full
	}
	return claimsOnly
}
