//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what TestDomainNames — 단일 사이트 빈 목록 / 도메인 모드 정렬 목록 검증

package yongol

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestDomainNames(t *testing.T) {
	cases := []struct {
		name string
		fs   *Fullstack
		want string // comma-joined expected names (sorted)
	}{
		{"single-site", &Fullstack{Manifest: nil}, ""},
		{
			"sorted",
			&Fullstack{Manifest: &manifest.ProjectConfig{
				Domains: map[string]manifest.DomainConfig{
					"public": {}, "admin": {}, "api": {},
				},
			}},
			"admin,api,public",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := strings.Join(tc.fs.DomainNames(), ",")
			if got != tc.want {
				t.Fatalf("DomainNames() = %q, want %q", got, tc.want)
			}
		})
	}
}
