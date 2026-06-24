//ff:func feature=validate type=test control=sequence
//ff:what TestStepRun — 도메인 모드 step.run 디스패치(single/once/merged/per-view/default) 검증

package validate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStepRun(t *testing.T) {
	t.Run("single-site runs once", func(t *testing.T) {
		var calls int
		s := step{
			Name:  "openapi",
			Kinds: []yongol.SSOTKind{yongol.KindOpenAPI},
			Run:   func(*yongol.Fullstack) []diagnostic.Diagnostic { calls++; return nil },
		}
		s.run(&yongol.Fullstack{OpenAPIDoc: &openapi3.T{}})
		if calls != 1 {
			t.Fatalf("single-site: want 1 call, got %d", calls)
		}
	})

	t.Run("DomainAware runs once on full fs", func(t *testing.T) {
		var calls int
		var sawDoc bool
		s := step{
			Name:        "domain_security",
			Kinds:       []yongol.SSOTKind{yongol.KindConfig, yongol.KindOpenAPI},
			DomainAware: true,
			Run: func(fs *yongol.Fullstack) []diagnostic.Diagnostic {
				calls++
				sawDoc = fs.OpenAPIDoc != nil
				return nil
			},
		}
		s.run(domainedFS())
		if calls != 1 {
			t.Fatalf("DomainAware: want 1 call, got %d", calls)
		}
		if sawDoc {
			t.Errorf("DomainAware: expected nil singular OpenAPIDoc (full fs)")
		}
	})

	t.Run("DomainMerged runs once on merged view", func(t *testing.T) {
		var calls int
		var sawDoc bool
		s := step{
			Name:         "openapi_manifest",
			Kinds:        []yongol.SSOTKind{yongol.KindOpenAPI, yongol.KindConfig},
			DomainMerged: true,
			Run: func(fs *yongol.Fullstack) []diagnostic.Diagnostic {
				calls++
				sawDoc = fs.OpenAPIDoc != nil
				return nil
			},
		}
		s.run(domainedFS())
		if calls != 1 {
			t.Fatalf("DomainMerged: want 1 call, got %d", calls)
		}
		if !sawDoc {
			t.Errorf("DomainMerged: expected merged OpenAPIDoc on the view")
		}
	})

	t.Run("per-view runs once per domain and aggregates", func(t *testing.T) {
		s := step{
			Name:  "openapi",
			Kinds: []yongol.SSOTKind{yongol.KindOpenAPI},
			Run: func(fs *yongol.Fullstack) []diagnostic.Diagnostic {
				if fs.OpenAPIDoc == nil {
					t.Errorf("per-view: expected per-domain OpenAPIDoc set")
				}
				return []diagnostic.Diagnostic{{Level: diagnostic.LevelWarning}}
			},
		}
		diags := s.run(domainedFS())
		if len(diags) != 2 {
			t.Fatalf("per-view: want 2 aggregated diags (one per domain), got %d", len(diags))
		}
	})

	t.Run("non per-domain step runs once in domain mode", func(t *testing.T) {
		var calls int
		s := step{
			Name:  "ddl",
			Kinds: []yongol.SSOTKind{yongol.KindDDL},
			Run:   func(*yongol.Fullstack) []diagnostic.Diagnostic { calls++; return nil },
		}
		s.run(domainedFS())
		if calls != 1 {
			t.Fatalf("default: want 1 call, got %d", calls)
		}
	})
}
