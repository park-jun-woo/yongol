//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-structural
//ff:what V-05 — components 내 토큰 참조 {group.token}이 같은 파일 내 실제 토큰으로 resolve 검증
package design

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var tokenRefRe = regexp.MustCompile(`\{([^}]+)\}`)

// v05TokenRefResolve validates that {group.token} references in component props
// resolve to actual tokens defined in the same DESIGN.md file.
func v05TokenRefResolve(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for compName, comp := range fs.DesignSpec.Components {
		diags = append(diags, checkPropRefs(fs, compName, comp.Props)...)
	}
	return diags
}
