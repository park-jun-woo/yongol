//ff:func feature=validate type=util control=iteration dimension=1 topic=design-structural
//ff:what checkPropRefs — component prop 값에서 {group.token} 참조를 찾아 resolve 여부 검사
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// checkPropRefs checks all {group.token} references in a single component's props
// and returns diagnostics for any unresolved references.
func checkPropRefs(fs *yongol.Fullstack, compName string, props map[string]string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for propName, propVal := range props {
		diags = append(diags, checkSinglePropRefs(fs, compName, propName, propVal)...)
	}
	return diags
}
