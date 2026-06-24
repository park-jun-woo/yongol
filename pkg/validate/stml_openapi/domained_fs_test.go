//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what domainedFS — frontend-ON 멀티 도메인 Fullstack 생성 (도메인별 doc/page)

package stml_openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// domainedFS builds a frontend-ON multi-domain Fullstack from per-domain docs
// and pages keyed by domain name.
func domainedFS(docs map[string]*openapi3.T, pages map[string][]stml.PageSpec) *yongol.Fullstack {
	domains := make(map[string]manifest.DomainConfig, len(docs))
	for name := range docs {
		domains[name] = manifest.DomainConfig{}
	}
	return &yongol.Fullstack{
		SpecsDir:          "/tmp/test-domained",
		Manifest:          &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}, Domains: domains},
		DomainOpenAPIDocs: docs,
		DomainSTMLPages:   pages,
	}
}
