//ff:func feature=stml-gen type=generator control=sequence
//ff:what 폼 필드(input/component)의 JSX를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFieldJSX generates JSX for a form field. idPrefix scopes the DOM
// id/htmlFor to the owning form (operationId-derived) so two forms on one
// page that declare a same-named field do not collide on a single id
// (BUG-127). The register/zod/formState.errors keys stay f.Name — those are
// the OpenAPI request-body contract; only the DOM id is form-scoped. An
// empty idPrefix leaves the bare field name (direct unit calls).
func renderFieldJSX(f stmlparser.FieldBind, formName, idPrefix string, indent int) string {
	ind := indentStr(indent)

	// data-component field — no label wrapper
	if strings.HasPrefix(f.Tag, "data-component:") {
		comp := strings.TrimPrefix(f.Tag, "data-component:")
		return fmt.Sprintf("%s<%s {...%s.register('%s')} />", ind, comp, formName, f.Name)
	}

	domID := f.Name
	if idPrefix != "" {
		domID = idPrefix + "-" + f.Name
	}

	var attrs []string
	attrs = append(attrs, fmt.Sprintf(`id="%s"`, domID))
	if f.Type != "" {
		attrs = append(attrs, fmt.Sprintf(`type="%s"`, f.Type))
	}
	if f.Placeholder != "" {
		attrs = append(attrs, fmt.Sprintf(`placeholder="%s"`, f.Placeholder))
	}
	if f.ClassName != "" {
		attrs = append(attrs, fmt.Sprintf(`className="%s"`, f.ClassName))
	}

	reg := fmt.Sprintf("{...%s.register('%s'", formName, f.Name)
	if f.Type == "number" {
		reg += ", { valueAsNumber: true }"
	}
	reg += ")}"
	if f.Type == "file" {
		reg = "{...(" + formName + ".register('" + f.Name + "') as any)}"
	}

	attrStr := " " + strings.Join(attrs, " ")
	label := toLabel(f.Name)

	var lines []string
	lines = append(lines, fmt.Sprintf(`%s<div className="space-y-1">`, ind))
	lines = append(lines, fmt.Sprintf(`%s  <label htmlFor="%s" className="text-sm font-medium">%s</label>`, ind, domID, label))
	lines = append(lines, fmt.Sprintf("%s  <Input%s %s />", ind, attrStr, reg))
	lines = append(lines, fmt.Sprintf(`%s  {%s.formState.errors.%s && (`, ind, formName, f.Name))
	lines = append(lines, fmt.Sprintf(`%s    <p className="text-sm text-destructive">{%s.formState.errors.%s?.message}</p>`, ind, formName, f.Name))
	lines = append(lines, fmt.Sprintf(`%s  )}`, ind))
	lines = append(lines, fmt.Sprintf("%s</div>", ind))
	return strings.Join(lines, "\n")
}
