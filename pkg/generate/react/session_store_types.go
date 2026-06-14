//ff:func feature=gen-react type=generator control=sequence
//ff:what sessionStoreTypes — bearer 세션 store 의 AuthState interface 선언 (hasRefresh 여부로 refresh 필드/인자 포함)

package react

// sessionStoreTypes returns the AuthState interface of the bearer session
// store. hasRefresh selects the shape: with the refresh token field and
// setAuth's optional second argument (resolveHasRefresh true), or the
// token-only reduction that drops the dead refresh surface (BUG-135).
func sessionStoreTypes(hasRefresh bool) string {
	const withRefresh = `
interface AuthState {
  token: string | null
  refresh: string | null
  setAuth: (token?: string | null, refresh?: string | null) => void
  clear: () => void
}

`
	const tokenOnly = `
interface AuthState {
  token: string | null
  setAuth: (token?: string | null) => void
  clear: () => void
}

`
	if hasRefresh {
		return withRefresh
	}
	return tokenOnly
}
