//ff:func feature=validate type=test control=sequence
//ff:what domainedFS — 2개 도메인(public/admin) doc 을 가진 멀티 도메인 Fullstack 픽스처

package validate

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// domainedFS builds a multi-domain Fullstack (IsDomained()==true) with two
// per-domain OpenAPI docs and a nil singular OpenAPIDoc, mirroring real
// domain-mode parse output.
func domainedFS() *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				"admin":  {},
				"public": {},
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{
			"public": {Paths: openapi3.NewPaths()},
			"admin":  {Paths: openapi3.NewPaths()},
		},
	}
}
