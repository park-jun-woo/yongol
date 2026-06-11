//ff:func feature=gen-react type=generator control=sequence
//ff:what claimsStoreTypes — claims store 의 AuthState interface 선언 (tokens 여부로 전체형/축소형)

package react

// claimsStoreTypes returns the AuthState interface of the claims store:
// the full bearer shape (token/refresh/setAuth + claims/setClaim) or the
// cookie-mode claims-only reduction (plans/stml/sitemap Phase005).
func claimsStoreTypes(tokens bool) string {
	const full = `
interface AuthState {
  token: string | null
  refresh: string | null
  claims: Record<string, string>
  setAuth: (token?: string | null, refresh?: string | null) => void
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
	if tokens {
		return full
	}
	return claimsOnly
}
