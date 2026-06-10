//ff:func feature=gen-react type=generator control=sequence
//ff:what writeRefreshFlow — single-flight 401→refresh→1회 재시도 흐름(withAuthRetry)을 api.ts 에 방출

package react

import (
	"fmt"
	"strings"
)

// writeRefreshFlow appends the bearer refresh flow to api.ts. The Bearer
// middleware (writeAuthzMiddleware with refresh) injects only; the 401
// semantics live here, at the operation-wrapper level, because the
// openapi-fetch ^0.13 middleware cannot retry the original request from
// onResponse (the request body is already consumed; see Phase004 사전 확정).
//
//   - single-flight: refreshInFlight is a module-level shared Promise so
//     concurrent 401s trigger exactly one refresh network call — duplicate
//     refreshes collide with server-side rotation/reuse detection (XNA-90)
//     and would force-close the session.
//   - retry once: withAuthRetry re-invokes the operation closure after a
//     successful refresh; a second 401 gives up. Refresh failure (or a
//     missing stored refresh token) clears the store and goes to /login.
func writeRefreshFlow(b *strings.Builder, rp *refreshPlan) {
	call := fmt.Sprintf("client.%s('%s', {})", strings.ToUpper(rp.method), rp.path)
	if rp.bodyKey != "" {
		call = fmt.Sprintf("client.%s('%s', { body: { %s: refresh } as any })", strings.ToUpper(rp.method), rp.path, rp.bodyKey)
	}
	b.WriteString(fmt.Sprintf(`
// 401 -> refresh (%s) -> retry once. Single-flight: concurrent 401s share
// the in-flight refresh Promise so the refresh endpoint — which rotates the
// token server-side — is hit exactly once per expiry.
let refreshInFlight: Promise<boolean> | null = null

function refreshSession(): Promise<boolean> {
  if (!refreshInFlight) {
    refreshInFlight = requestRefresh().finally(() => {
      refreshInFlight = null
    })
  }
  return refreshInFlight
}

async function requestRefresh(): Promise<boolean> {
  const refresh = useAuthStore.getState().refresh
  if (!refresh) return false
  const r = await %s
  const data = r.data as Record<string, any> | undefined
  if (!r.response.ok || !data) return false
  useAuthStore.getState().setAuth(data['%s'], data['%s'])
  return true
}

function clearSessionAndLogin() {
  useAuthStore.getState().clear()
  window.location.href = '/login'
}

async function withAuthRetry<T extends { response: Response }>(call: () => Promise<T>): Promise<T> {
  const first = await call()
  if (first.response.status !== 401) return first
  if (!(await refreshSession())) {
    clearSessionAndLogin()
    return first
  }
  const retried = await call()
  if (retried.response.status === 401) {
    clearSessionAndLogin()
  }
  return retried
}
`, rp.opID, call, rp.tokenField, rp.refreshField))
}
