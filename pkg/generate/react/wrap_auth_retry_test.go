//ff:func feature=gen-react type=test control=sequence
//ff:what wrapAuthRetry — withRetry 시 withAuthRetry 클로저로 감싸고, 아니면 원식 그대로 검증

package react

import "testing"

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
}
