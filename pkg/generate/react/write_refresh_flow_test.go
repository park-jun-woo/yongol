//ff:func feature=gen-react type=test control=sequence
//ff:what writeRefreshFlow — single-flight 401→refresh→재시도 흐름 방출, bodyKey 유무에 따른 호출식 분기 검증

package react

import (
	"strings"
	"testing"
)

func TestWriteRefreshFlow(t *testing.T) {
	// bodyKey set -> refresh call sends { <key>: refresh } body
	var b strings.Builder
	writeRefreshFlow(&b, &refreshPlan{
		opID: "Refresh", method: "post", path: "/auth/refresh",
		tokenField: "access_token", refreshField: "refresh_token", bodyKey: "refresh_token",
	})
	out := b.String()

	// single-flight shared promise + retry-once wrapper
	assertContains(t, out, "let refreshInFlight: Promise<boolean> | null = null")
	assertContains(t, out, "async function withAuthRetry")
	// method upper-cased, body carries the stored refresh token
	assertContains(t, out, `client.POST('/auth/refresh', { body: { refresh_token: refresh } as any })`)
	// 2xx fields committed back into the store
	assertContains(t, out, "useAuthStore.getState().setAuth(data['access_token'], data['refresh_token'])")
	// refresh failure path clears + redirects
	assertContains(t, out, "window.location.href = '/login'")

	// bodyKey empty -> refresh call sends no body (e.g. cookie-carried refresh)
	var b2 strings.Builder
	writeRefreshFlow(&b2, &refreshPlan{
		opID: "Refresh", method: "post", path: "/auth/refresh",
		tokenField: "access_token", refreshField: "refresh_token", bodyKey: "",
	})
	out2 := b2.String()
	assertContains(t, out2, `client.POST('/auth/refresh', {})`)
	assertNotContains(t, out2, "body:")
}
