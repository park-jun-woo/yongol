//ff:func feature=gen-react type=generator control=sequence
//ff:what claimsStoreTypes — claims store 의 AuthState interface 선언 (tokens 여부로 전체형/축소형)

package react

// claimsStoreTypes returns the AuthState interface of the claims store:
// the full bearer shape (token/setAuth + claims/setClaim) or the cookie-mode
// claims-only reduction (plans/stml/sitemap Phase005). In the bearer shape
// hasRefresh (resolveHasRefresh) gates the refresh token field and setAuth's
// optional second argument — false drops the dead refresh surface (BUG-135).
func claimsStoreTypes(tokens, hasRefresh bool) string {
	const fullRefresh = `
interface AuthState {
  token: string | null
  refresh: string | null
  claims: Record<string, string>
  setAuth: (token?: string | null, refresh?: string | null) => void
  setClaim: (name: string, value: string) => void
  clear: () => void
}

`
	const fullNoRefresh = `
interface AuthState {
  token: string | null
  claims: Record<string, string>
  setAuth: (token?: string | null) => void
  setClaim: (name: string, value: string) => void
  clear: () => void
}

`
	const claimsOnly = `
interface AuthState {
  claims: Record<string, string>
  setClaim: (name: string, value: string) => void
  clear: () => void
}

`
	if !tokens {
		return claimsOnly
	}
	if hasRefresh {
		return fullRefresh
	}
	return fullNoRefresh
}
