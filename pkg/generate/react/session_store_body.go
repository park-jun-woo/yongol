//ff:func feature=gen-react type=generator control=sequence
//ff:what sessionStoreBody — bearer 세션 store 상태 객체 본문 라인들 (hasRefresh 여부로 refresh 필드/커밋 포함)

package react

// sessionStoreBody returns the state-object body lines of the bearer session
// store (rendered at both persist and memory depths via indentLines).
// hasRefresh selects the shape: the full token/refresh contract
// (resolveHasRefresh true) — byte-identical to the pre-BUG-135 output — or
// the token-only reduction that drops the dead refresh field, its setAuth
// argument and clear's reset of it.
func sessionStoreBody(hasRefresh bool) []string {
	if hasRefresh {
		return []string{
			"token: null,",
			"refresh: null,",
			"setAuth: (token, refresh) =>",
			"  set((state) => ({",
			"    token: token ?? null,",
			"    refresh: refresh === undefined ? state.refresh : refresh,",
			"  })),",
			"clear: () => set({ token: null, refresh: null }),",
		}
	}
	return []string{
		"token: null,",
		"setAuth: (token) => set({ token: token ?? null }),",
		"clear: () => set({ token: null }),",
	}
}
