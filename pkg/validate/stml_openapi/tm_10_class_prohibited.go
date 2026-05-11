//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-10 — STML 요소에 class 속성 직접 사용 금지 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm10ClassProhibited scans all STML pages and emits TM-10 ERROR for any
// element that has a non-empty ClassName. Designer custom styles must use
// <!-- @override class="..." --> comments instead.
func tm10ClassProhibited(pages []stml.PageSpec) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, page := range pages {
		for _, f := range page.Fetches {
			diags = append(diags, checkFetchClass(f, page.FileName)...)
		}
		for _, a := range page.Actions {
			diags = append(diags, checkActionClass(a, page.FileName)...)
		}
		for _, c := range page.Children {
			diags = append(diags, checkChildClass(c, page.FileName)...)
		}
	}
	return diags
}

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

// checkStateClass checks a StateBind and its descendants for class usage.
func checkStateClass(sb stml.StateBind, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "StateBind", sb.Condition, sb.ClassName)...)
	for _, c := range sb.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	return diags
}

// checkChildClass dispatches class checks on a ChildNode.
func checkChildClass(cn stml.ChildNode, file string) []diagnostic.Diagnostic {
	switch cn.Kind {
	case "static":
		if cn.Static != nil {
			return checkStaticClass(*cn.Static, file)
		}
	case "fetch":
		if cn.Fetch != nil {
			return checkFetchClass(*cn.Fetch, file)
		}
	case "action":
		if cn.Action != nil {
			return checkActionClass(*cn.Action, file)
		}
	case "each":
		if cn.Each != nil {
			return checkEachClass(*cn.Each, file)
		}
	case "state":
		if cn.State != nil {
			return checkStateClass(*cn.State, file)
		}
	case "component":
		if cn.Component != nil {
			return emitIfClass(file, "ComponentRef", cn.Component.Name, cn.Component.ClassName)
		}
	case "bind":
		if cn.Bind != nil {
			return emitIfClass(file, "FieldBind", cn.Bind.Name, cn.Bind.ClassName)
		}
	}
	return nil
}

// checkStaticClass checks a StaticElement and its descendants for class usage.
func checkStaticClass(se stml.StaticElement, file string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, emitIfClass(file, "StaticElement", se.Tag, se.ClassName)...)
	for _, c := range se.Children {
		diags = append(diags, checkChildClass(c, file)...)
	}
	return diags
}

// emitIfClass returns a TM-10 diagnostic if className is non-empty.
func emitIfClass(file, elemType, elemID, className string) []diagnostic.Diagnostic {
	if className == "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-10] class attribute %q on %s %q is prohibited in STML; use <!-- @override class=\"...\" --> comment instead", className, elemType, elemID),
		Advice:  "Remove the class attribute from the STML element. To override DESIGN.md styling, place <!-- @override class=\"...\" --> as a comment before the element",
	}}
}
