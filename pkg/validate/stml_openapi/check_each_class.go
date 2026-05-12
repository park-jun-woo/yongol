//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what checkEachClass — EachBlock 와 하위 요소의 class 사용 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkEachClass checks an EachBlock and its descendants for class usage.
func checkEachClass(eb stml.EachBlock, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "EachBlock", eb.Field, eb.ClassName)...)
	if eb.ItemClassName != "" {
		diags = append(diags, emitIfClass(file, "EachBlock item", eb.Field, eb.ItemClassName)...)
	}
	for _, b := range eb.Binds {
		diags = append(diags, emitIfClass(file, "FieldBind", b.Name, b.ClassName)...)
	}
	for _, s := range eb.States {
		diags = append(diags, checkStateClass(s, file)...)
	}
	for _, comp := range eb.Components {
		diags = append(diags, emitIfClass(file, "ComponentRef", comp.Name, comp.ClassName)...)
	}
	for _, c := range eb.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	return diags
}
