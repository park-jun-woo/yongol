//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestClaimsSinkName — auth.claims.<name> 판정의 양성·음성(접두사/빈 이름/비식별자) 검증

package stml

import "testing"

func TestClaimsSinkName(t *testing.T) {
	cases := []struct {
		sink string
		name string
		ok   bool
	}{
		{"auth.claims.role", "role", true},
		{"auth.claims.user_role", "user_role", true},
		{"auth.claims._r2", "_r2", true},
		{"auth.claims.Role9", "Role9", true},
		{"auth.claims.", "", false},
		{"auth.claims.9role", "", false},
		{"auth.claims.ro-le", "", false},
		{"auth.claims.ro.le", "", false},
		{"auth.token", "", false},
		{"auth.refresh", "", false},
		{"claims.role", "", false},
		{"", "", false},
	}
	for _, c := range cases {
		name, ok := ClaimsSinkName(c.sink)
		if name != c.name || ok != c.ok {
			t.Errorf("ClaimsSinkName(%q) = (%q, %v), want (%q, %v)", c.sink, name, ok, c.name, c.ok)
		}
	}
}
