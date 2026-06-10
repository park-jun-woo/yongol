//ff:func feature=manifest type=test control=sequence
//ff:what ResolvedStore — 빈 값/nil 은 localStorage, 명시 값은 그대로

package manifest

import (
	"testing"
)

func TestResolvedStore(t *testing.T) {
	if got := (&FrontendAuth{}).ResolvedStore(); got != "localStorage" {
		t.Errorf("empty store → %q, want localStorage", got)
	}
	if got := (&FrontendAuth{Store: "memory"}).ResolvedStore(); got != "memory" {
		t.Errorf("explicit memory → %q", got)
	}
	if got := (&FrontendAuth{Store: "localStorage"}).ResolvedStore(); got != "localStorage" {
		t.Errorf("explicit localStorage → %q", got)
	}
	var nilAuth *FrontendAuth
	if got := nilAuth.ResolvedStore(); got != "localStorage" {
		t.Errorf("nil auth → %q, want localStorage", got)
	}
}
