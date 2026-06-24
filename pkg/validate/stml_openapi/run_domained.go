//ff:func feature=validate type=rule control=sequence dimension=1 topic=stml-openapi
//ff:what runDomained — 도메인 모드 STML<->OpenAPI 커버리지: XMO-10 스킵, XMO-11/12 를 전체 도메인 집계로 평가

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runDomained runs the domain-agnostic STML coverage rules for a multi-domain
// project. XMO-10 (single-site operationId consumption) is the responsibility
// of domain_security's XMO-20/21/22 here and is skipped; XMO-11 (frontend ON
// with zero pages) and XMO-12 (stale no-front tag) remain meaningful and are
// evaluated across every domain via the aggregate variants.
func runDomained(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xmo11NoStmlAll(fs)...)
	diags = append(diags, xmo12NoFrontConsumedAll(fs)...)
	return diags
}
