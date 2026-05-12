//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what checkFetchClass — FetchBlock 와 하위 요소의 class 사용 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkFetchClass checks a FetchBlock and its descendants for class usage.
func checkFetchClass(fb stml.FetchBlock, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "FetchBlock", fb.OperationID, fb.ClassName)...)
	for _, b := range fb.Binds {
		diags = append(diags, emitIfClass(file, "FieldBind", b.Name, b.ClassName)...)
	}
	for _, e := range fb.Eaches {
		diags = append(diags, checkEachClass(e, file)...)
	}
	for _, s := range fb.States {
		diags = append(diags, checkStateClass(s, file)...)
	}
	for _, comp := range fb.Components {
		diags = append(diags, emitIfClass(file, "ComponentRef", comp.Name, comp.ClassName)...)
	}
	for _, c := range fb.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	for _, nf := range fb.NestedFetches {
		diags = append(diags, checkFetchClass(nf, file)...)
	}
	return diags
}
