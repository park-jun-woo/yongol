//ff:func feature=gen-react type=test control=sequence
//ff:what wrapAuthRetry — withRetry 래핑 여부 + throw 승격 .then 이 retry 클로저 바깥임을 고정 검증

package react

import (
	"strings"
	"testing"
)

func TestWrapAuthRetry(t *testing.T) {
	call := "client.GET('/items', {})"

	// withRetry false -> call returned unchanged
	if got := wrapAuthRetry(call, false); got != call {
		t.Errorf("withRetry=false = %q, want unchanged %q", got, call)
	}

	// withRetry true -> wrapped in withAuthRetry(() => ...)
	want := "withAuthRetry(() => " + call + ")"
	if got := wrapAuthRetry(call, true); got != want {
		t.Errorf("withRetry=true = %q, want %q", got, want)
	}

	// BUG-113 order guarantee: the error-promoting .then attaches to the
	// withAuthRetry result, never inside the retried closure — the
	// 401→refresh→retry flow completes first, then the final result throws.
	var b strings.Builder
	writeApiClientEntry(&b, endpoint{method: "get", path: "/items", opID: "ListItems"}, true)
	entry := b.String()
	assertContains(t, entry,
		"withAuthRetry(() => client.GET('/items', { params: { query: (args ?? {}) as any } })).then(r => { const d = r.data; const e = r.error; if (e !== undefined) throw e; return d as Res<'ListItems'> })")
	assertNotContains(t, entry, "throw e; return d as Res<'ListItems'> }))")
}
