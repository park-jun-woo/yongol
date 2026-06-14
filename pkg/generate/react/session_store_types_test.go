//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what sessionStoreTypes — hasRefresh true/false 분기별 AuthState interface 셰이프 검증

package react

import (
	"strings"
	"testing"
)

func TestSessionStoreTypes(t *testing.T) {
	withRefresh := sessionStoreTypes(true)
	if !strings.Contains(withRefresh, "refresh: string | null") {
		t.Errorf("hasRefresh=true: missing refresh field\n%s", withRefresh)
	}
	if !strings.Contains(withRefresh, "setAuth: (token?: string | null, refresh?: string | null) => void") {
		t.Errorf("hasRefresh=true: missing two-arg setAuth signature\n%s", withRefresh)
	}

	tokenOnly := sessionStoreTypes(false)
	if strings.Contains(tokenOnly, "refresh") {
		t.Errorf("hasRefresh=false: refresh surface must be dropped\n%s", tokenOnly)
	}
	if !strings.Contains(tokenOnly, "setAuth: (token?: string | null) => void") {
		t.Errorf("hasRefresh=false: missing one-arg setAuth signature\n%s", tokenOnly)
	}

	// both shapes declare the AuthState interface and a clear() method
	for _, out := range []string{withRefresh, tokenOnly} {
		if !strings.Contains(out, "interface AuthState {") {
			t.Errorf("missing AuthState interface decl\n%s", out)
		}
		if !strings.Contains(out, "clear: () => void") {
			t.Errorf("missing clear method\n%s", out)
		}
	}
}
