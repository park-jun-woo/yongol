//ff:func feature=validate type=accessor control=iteration dimension=1
//ff:what step.usesPerDomainDoc — OpenAPI/STML 게이트 step(=도메인별 단수 필드 사용) 판정
package validate

import "github.com/park-jun-woo/yongol/pkg/yongol"

// usesPerDomainDoc reports whether the step gates on a per-domain singular SSOT
// (OpenAPI or STML) and therefore needs per-domain-view scoping in domain mode.
func (s step) usesPerDomainDoc() bool {
	for _, k := range s.Kinds {
		if k == yongol.KindOpenAPI || k == yongol.KindSTML {
			return true
		}
	}
	return false
}
