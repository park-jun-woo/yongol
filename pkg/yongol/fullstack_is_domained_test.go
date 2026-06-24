//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what TestIsDomained — manifest domains 선언 유무에 따른 IsDomained 분기 검증

package yongol

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestIsDomained(t *testing.T) {
	cases := []struct {
		name string
		fs   *Fullstack
		want bool
	}{
		{"nil-manifest", &Fullstack{Manifest: nil}, false},
		{"empty-domains", &Fullstack{Manifest: &manifest.ProjectConfig{}}, false},
		{
			"with-domains",
			&Fullstack{Manifest: &manifest.ProjectConfig{
				Domains: map[string]manifest.DomainConfig{"public": {}},
			}},
			true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fs.IsDomained(); got != tc.want {
				t.Fatalf("IsDomained() = %v, want %v", got, tc.want)
			}
		})
	}
}
