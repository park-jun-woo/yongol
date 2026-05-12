//ff:type feature=validate type=model
//ff:what domainDoc — 도메인별 OpenAPI 문서와 설정을 담는 구조체
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// domainDoc holds a loaded OpenAPI document and its domain name.
type domainDoc struct {
	Name string
	Doc  *openapi3.T
	Cfg  manifest.DomainConfig
}
