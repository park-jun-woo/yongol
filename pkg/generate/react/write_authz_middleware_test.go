//ff:func feature=gen-react type=test control=sequence
//ff:what writeAuthzMiddleware — withRefresh 분기별 onRequest/onResponse 방출 검증

package react

import (
	"strings"
	"testing"
)

func TestWriteAuthzMiddleware(t *testing.T) {
	t.Run("no refresh", func(t *testing.T) {
		var b strings.Builder
		writeAuthzMiddleware(&b, false)
		out := b.String()

		assertContains(t, out, "async onRequest({ request })")
		assertContains(t, out, "request.headers.set('Authorization', `Bearer ${token}`)")
		assertContains(t, out, "request.headers.has('Authorization')")
		assertContains(t, out, "{ request, response }")
	})

	t.Run("with refresh", func(t *testing.T) {
		var b strings.Builder
		writeAuthzMiddleware(&b, true)
		out := b.String()

		assertContains(t, out, "async onRequest({ request })")
		assertContains(t, out, "request.headers.set('Authorization', `Bearer ${token}`)")
		assertNotContains(t, out, "onResponse")
	})
}
