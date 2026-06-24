//ff:func feature=orchestrator type=test control=sequence
//ff:what TestMergedOpenAPIView — 도메인 doc 합집합(Paths/Components 유니온, dedup, 수신자 불변) 검증

package yongol

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestMergedOpenAPIView(t *testing.T) {
	t.Run("unions paths and components, dedups schemes, preserves receiver", func(t *testing.T) {
		pub := &openapi3.T{
			Paths: openapi3.NewPaths(),
			Components: &openapi3.Components{
				Schemas:         openapi3.Schemas{"Post": &openapi3.SchemaRef{}},
				SecuritySchemes: openapi3.SecuritySchemes{"bearerAuth": &openapi3.SecuritySchemeRef{}},
			},
		}
		pub.Paths.Set("/posts", &openapi3.PathItem{})
		adm := &openapi3.T{
			Paths: openapi3.NewPaths(),
			Components: &openapi3.Components{
				Schemas:         openapi3.Schemas{"Report": &openapi3.SchemaRef{}},
				SecuritySchemes: openapi3.SecuritySchemes{"bearerAuth": &openapi3.SecuritySchemeRef{}},
			},
		}
		adm.Paths.Set("/reports", &openapi3.PathItem{})

		fs := &Fullstack{
			Manifest:          &manifest.ProjectConfig{Domains: map[string]manifest.DomainConfig{"public": {}, "admin": {}}},
			DomainOpenAPIDocs: map[string]*openapi3.T{"public": pub, "admin": adm},
		}

		view := fs.MergedOpenAPIView()
		if view.OpenAPIDoc == nil {
			t.Fatal("merged OpenAPIDoc is nil")
		}
		if view.OpenAPIDoc.Paths.Find("/posts") == nil || view.OpenAPIDoc.Paths.Find("/reports") == nil {
			t.Errorf("merged paths missing a domain: got %v", view.OpenAPIDoc.Paths.Map())
		}
		if len(view.OpenAPIDoc.Components.Schemas) != 2 {
			t.Errorf("merged schemas = %d, want 2", len(view.OpenAPIDoc.Components.Schemas))
		}
		if len(view.OpenAPIDoc.Components.SecuritySchemes) != 1 {
			t.Errorf("merged security schemes = %d, want 1 (deduped)", len(view.OpenAPIDoc.Components.SecuritySchemes))
		}
		if fs.OpenAPIDoc != nil {
			t.Errorf("MergedOpenAPIView mutated receiver OpenAPIDoc")
		}
	})

	t.Run("nil domain docs are skipped", func(t *testing.T) {
		fs := &Fullstack{
			Manifest:          &manifest.ProjectConfig{Domains: map[string]manifest.DomainConfig{"public": {}, "admin": {}}},
			DomainOpenAPIDocs: map[string]*openapi3.T{"public": nil, "admin": nil},
		}
		view := fs.MergedOpenAPIView()
		if view.OpenAPIDoc == nil || view.OpenAPIDoc.Paths == nil {
			t.Fatal("merged doc/paths must be non-nil even when all domain docs are nil")
		}
		if view.OpenAPIDoc.Paths.Len() != 0 {
			t.Errorf("want 0 merged paths, got %d", view.OpenAPIDoc.Paths.Len())
		}
	})
}
