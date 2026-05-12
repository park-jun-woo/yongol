//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what checkStaticClass — StaticElement 와 하위 요소의 class 사용 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkStaticClass checks a StaticElement and its descendants for class usage.
func checkStaticClass(se stml.StaticElement, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "StaticElement", se.Tag, se.ClassName)...)
	for _, c := range se.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	return diags
}
