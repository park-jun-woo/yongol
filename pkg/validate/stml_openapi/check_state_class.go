//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what checkStateClass — StateBind 와 하위 요소의 class 사용 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkStateClass checks a StateBind and its descendants for class usage.
func checkStateClass(sb stml.StateBind, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "StateBind", sb.Condition, sb.ClassName)...)
	for _, c := range sb.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	return diags
}
