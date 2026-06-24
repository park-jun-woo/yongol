//ff:func feature=projectconfig type=test control=iteration dimension=1
//ff:what DomainConfig.ResolvedAuthMode — override vs backend 상속 분기 검증

package manifest

import "testing"

func TestDomainConfigResolvedAuthMode(t *testing.T) {
	cases := []struct {
		name     string
		authMode string
		backend  string
		want     string
	}{
		{"override wins", "bearer", "cookie", "bearer"},
		{"inherit backend", "", "cookie", "cookie"},
		{"inherit empty backend", "", "", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			d := DomainConfig{AuthMode: c.authMode}
			if got := d.ResolvedAuthMode(c.backend); got != c.want {
				t.Errorf("ResolvedAuthMode(%q) with auth_mode=%q = %q, want %q", c.backend, c.authMode, got, c.want)
			}
		})
	}
}
