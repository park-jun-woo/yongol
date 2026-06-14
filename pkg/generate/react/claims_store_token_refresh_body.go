//ff:func feature=gen-react type=generator control=sequence
//ff:what claimsStoreTokenRefreshBody — bearer+refresh claims store 상태 객체 본문 라인들 (token/refresh/claims + setAuth/setClaim/clear)

package react

// claimsStoreTokenRefreshBody returns the state-object body of the claims
// store in the bearer shape that carries the refresh token (resolveHasRefresh
// true). setAuth is byte-identical to writeSessionStore's — it never touches
// claims (zustand's shallow merge preserves them across a 401 refresh) — and
// clear() resets claims together with the tokens (BUG-135). Extracted from
// claimsStoreBody to keep that branch within the per-block line limit.
func claimsStoreTokenRefreshBody() []string {
	return []string{
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
}
