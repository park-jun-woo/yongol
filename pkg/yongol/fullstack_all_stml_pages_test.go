//ff:func feature=orchestrator type=test control=sequence
//ff:what TestAllSTMLPages — 단일 사이트 패스스루 / 도메인 모드 정렬 평탄화 검증

package yongol

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAllSTMLPages(t *testing.T) {
	t.Run("single-site", func(t *testing.T) {
		pages := []stml.PageSpec{{Name: "home"}}
		fs := &Fullstack{STMLPages: pages}
		got := fs.AllSTMLPages()
		if len(got) != 1 || got[0].Name != "home" {
			t.Fatalf("got = %v, want [home]", got)
		}
	})

	t.Run("domained-flatten-sorted", func(t *testing.T) {
		fs := &Fullstack{
			Manifest: &manifest.ProjectConfig{Domains: map[string]manifest.DomainConfig{
				"public": {}, "admin": {},
			}},
			DomainSTMLPages: map[string][]stml.PageSpec{
				"public": {{Name: "p1"}, {Name: "p2"}},
				"admin":  {{Name: "a1"}},
			},
		}
		got := fs.AllSTMLPages()
		// DomainNames order is sorted: admin, public.
		want := []string{"a1", "p1", "p2"}
		if len(got) != len(want) {
			t.Fatalf("got %v, want order %v", got, want)
		}
		for i, n := range want {
			if got[i].Name != n {
				t.Fatalf("got[%d].Name = %q, want %q (full %v)", i, got[i].Name, n, got)
			}
		}
	})
}
