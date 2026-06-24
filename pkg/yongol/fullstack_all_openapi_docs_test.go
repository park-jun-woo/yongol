//ff:func feature=orchestrator type=test control=sequence
//ff:what TestAllOpenAPIDocs — 단일 사이트 {"":doc} fallback / 도메인 모드 map 패스스루 검증

package yongol

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestAllOpenAPIDocs(t *testing.T) {
	single := &openapi3.T{}
	pub := &openapi3.T{}

	t.Run("single-site", func(t *testing.T) {
		fs := &Fullstack{OpenAPIDoc: single}
		got := fs.AllOpenAPIDocs()
		if len(got) != 1 {
			t.Fatalf("len = %d, want 1", len(got))
		}
		if got[""] != single {
			t.Fatalf("got[\"\"] = %p, want %p", got[""], single)
		}
	})

	t.Run("domained", func(t *testing.T) {
		domainMap := map[string]*openapi3.T{"public": pub}
		fs := &Fullstack{
			Manifest:          &manifest.ProjectConfig{Domains: map[string]manifest.DomainConfig{"public": {}}},
			DomainOpenAPIDocs: domainMap,
		}
		got := fs.AllOpenAPIDocs()
		if len(got) != 1 || got["public"] != pub {
			t.Fatalf("got = %v, want {public: %p}", got, pub)
		}
	})
}
