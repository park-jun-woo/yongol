//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what checkActionClass — ActionBlock 와 하위 요소의 class 사용 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkActionClass checks an ActionBlock and its descendants for class usage.
func checkActionClass(ab stml.ActionBlock, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "ActionBlock", ab.OperationID, ab.ClassName)...)
	for _, f := range ab.Fields {
		diags = append(diags, emitIfClass(file, "FieldBind", f.Name, f.ClassName)...)
	}
	for _, c := range ab.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	return diags
}
