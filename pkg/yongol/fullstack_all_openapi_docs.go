//ff:func feature=orchestrator type=accessor control=sequence
//ff:what AllOpenAPIDocs — 도메인 모드 시 도메인별 doc map, 단일 사이트 시 {"": doc}

package yongol

import "github.com/getkin/kin-openapi/openapi3"

// AllOpenAPIDocs returns every OpenAPI document keyed by domain name. In domain
// mode it returns the per-domain DomainOpenAPIDocs map directly; in single-site
// mode it returns the singular doc under the empty-string key so callers can
// iterate uniformly.
func (fs *Fullstack) AllOpenAPIDocs() map[string]*openapi3.T {
	if fs.IsDomained() {
		return fs.DomainOpenAPIDocs
	}
	return map[string]*openapi3.T{"": fs.OpenAPIDoc}
}
