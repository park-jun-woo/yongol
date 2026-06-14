//ff:func feature=gen-react type=test control=sequence
//ff:what sessionStoreBody — hasRefresh true/false 분기별 store 본문 라인 셰이프 검증

package react

import (
	"strings"
	"testing"
)

func TestSessionStoreBody(t *testing.T) {
	withRefresh := sessionStoreBody(true)
	joined := strings.Join(withRefresh, "\n")
	if !strings.Contains(joined, "refresh: null,") {
		t.Errorf("hasRefresh=true: missing refresh field\n%s", joined)
	}
	if !strings.Contains(joined, "setAuth: (token, refresh) =>") {
		t.Errorf("hasRefresh=true: missing two-arg setAuth\n%s", joined)
	}
	if !strings.Contains(joined, "clear: () => set({ token: null, refresh: null }),") {
		t.Errorf("hasRefresh=true: clear must reset refresh\n%s", joined)
	}

	tokenOnly := sessionStoreBody(false)
	joinedTO := strings.Join(tokenOnly, "\n")
	if strings.Contains(joinedTO, "refresh") {
		t.Errorf("hasRefresh=false: refresh field must be dropped\n%s", joinedTO)
	}
	if !strings.Contains(joinedTO, "setAuth: (token) => set({ token: token ?? null }),") {
		t.Errorf("hasRefresh=false: missing one-arg setAuth\n%s", joinedTO)
	}
	if len(tokenOnly) >= len(withRefresh) {
		t.Errorf("token-only body should be shorter: %d vs %d", len(tokenOnly), len(withRefresh))
	}
}
